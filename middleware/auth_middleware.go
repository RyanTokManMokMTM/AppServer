package middleware
import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
TODO - JWT Explanation
- HEADER
fields : encode method , type of token

{
	"alg":"HS256", //Hash 256 algorithm
	"typ":"JWT" //JSON WEB TOKEN type
}

- payload
field : information - iss(issuer),exp,sub etc...(registered )
iss (Issuer) - jwt簽發者
sub (Subject) - jwt所面向的用戶
aud (Audience) - 接收jwt的一方
exp (Expiration Time) - jwt的過期時間，這個過期時間必須要大於簽發時間
nbf (Not Before) - 定義在什麼時間之前，該jwt都是不可用的
iat (Issued At) - jwt的簽發時間
jti (JWT ID) - jwt的唯一身份標識，主要用來作為一次性token,從而迴避重放攻擊

{
	"sub":"server-api", //registered
	"exp": 15464664, // time in ms ,registered
	"name":"jackson" //public claims
	"email":"admin@admin.com" //public claims
}

JWT signature: HMACSHA256(encoder header .encode payload,secret key(server side))
signature : HMACSHA256(header.payload.key)
TODO - TOKEN store at Header field : Authorization : Bearer tokenString
*/

const (
	ISSUER string = "jackson.tmm"
	KEY string = "2zd2Y4YDwnBBqfMiFIUt8A8toecG1DRAsPwoWnEJCheXcr-h2FTvdo595Z8uL8_atA625OTnC7OocC7Rc_SCQQ=="
)

var (
	TokenExpired error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("")
	TokenMalformed error = errors.New("")
	TokenInvalid error = errors.New("")
)

type (
	//JWT Signature
	JWT struct {
		//Token from client
		SignKey []byte
	}

	//CustomClaims Payload
	CustomClaims struct{
		Name string  //Public info
		Email string  //Public info
		jwt.StandardClaims //jwt stander info header and payload
	}
)

//CreateToken -Generate token base on CustomClaims and BaseHash256
func (jwtObj *JWT) CreateToken(info CustomClaims) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	return token.SignedString(jwtObj.SignKey) //sign the token with our key
}

//ParseToken -Parse token
func (jwtObj *JWT) ParseToken(jwtToken string) (*CustomClaims ,error){
	token, err := jwt.ParseWithClaims(jwtToken, &CustomClaims{}, func(jwt *jwt.Token) (interface{}, error) {
		return jwtObj.SignKey, nil
	})

	if err != nil{
		if errType, ok := err.(*jwt.ValidationError);ok {
			if errType.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if errType.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if errType.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claim,ok := token.Claims.(*CustomClaims);ok && token.Valid{
		return claim,nil
	}
	return nil,fmt.Errorf("invalid token")
}

func NewJsonWebToken() *JWT{
	return &JWT{SignKey: []byte(KEY)}
}

func JWTAuth() gin.HandlerFunc{
	return func(ctx* gin.Context){
		authHeader := ctx.Request.Header.Get("authorization")
		if authHeader == ""{
			ctx.JSON(http.StatusOK,gin.H{
				"status":-1,
				"msg":"token is not found",
				"data":nil,
			})
			ctx.Abort()
			return
		}

		log.Printf("Token:%s",authHeader)
		jwt := NewJsonWebToken()

		//parse token into jwt object
		claimsInfo, err := jwt.ParseToken(authHeader)
		if err != nil {
			if err == TokenExpired{
				ctx.JSON(http.StatusOK,gin.H{
					"status":-1,
					"msg":"Token is expired",
					"data":nil,
				})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK,gin.H{
				"status":-1,
				"msg":err.Error(),
				"data":nil,
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims",claimsInfo)
	}
}
