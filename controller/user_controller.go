package controller

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	orm "music_api_server/database"
	mid "music_api_server/middleware"
	"music_api_server/model"
	"net/http"
	"os"
	"time"
)

var route *log.Logger
func init(){
	route = log.New(os.Stdout,"Route debug[DEBUG]",log.LstdFlags)
}

func UserLogin(ctx *gin.Context){
	//TODO -Info checking
	//TODO -User
	if orm.DB == nil{
		ctx.String(http.StatusInternalServerError,"server initialize apiError")
		ctx.Abort()
		return
	}

	var user model.UserInfo
	if err := ctx.BindJSON(&user) ;err != nil{
		ctx.JSON(http.StatusOK,gin.H{
			"status":-1,
			"msg":"Authorization failed",
			"data":nil,
		})
		ctx.Abort()
		return
	}else{
		var dbResult model.UserInfo
		if err := orm.DB.Where("username = ?","jackson").First(&dbResult);err != nil{
			if errors.Is(err.Error,gorm.ErrRecordNotFound){
				ctx.JSON(http.StatusOK,gin.H{
					"status":-1,
					"msg":"username or password incorrect",
					"data":nil,
				})
				ctx.Abort()
				return
			}else if err.Error != nil{
				ctx.JSON(http.StatusOK,gin.H{
					"status":-1,
					"msg":err.Error,
					"data":nil,
				})
				ctx.Abort()
				return
			}
		}
		if dbResult.Password != user.Password{
			ctx.JSON(http.StatusOK,gin.H{
				"state":-1,
				"msg":"user name or password incorrect",
				"data":nil,
			})
			ctx.Abort()
			return
		}

		//TODO -Token
		tokenGenerator(ctx,&dbResult)
		return
	}
}

func UserSignUp(ctx *gin.Context){
	//TODO -User SignUp
	//TODO -Info Checking
}

func UserSignOut(ctx *gin.Context){
	//TODO -SOMETHING
}



func tokenGenerator(ctx *gin.Context,info *model.UserInfo){
	token := mid.NewJsonWebToken()
	claims := mid.CustomClaims{
		Name:  info.Username,
		Email: info.Email,
		StandardClaims: jwt.StandardClaims{
			Issuer:"jackson.tmm",
			ExpiresAt: int64(time.Now().Unix() + 3600),
			NotBefore: int64(time.Now().Unix() - 1000),
			Subject: "User Authorization",
		},
	}

	jsonToken, err := token.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"status":-1,
			"msg":err.Error(),
			"data":nil,
		})
	}

	log.Printf("User Token is %v",jsonToken)

	ctx.JSON(http.StatusOK,gin.H{
		"status":0,
		"msg":"Logged in...",
		"data": map[string]interface{}{
			"user name":info.Username,
			"token":jsonToken,
		},
	})
	return
}