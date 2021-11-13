package user_controller

import (
	"github.com/gin-gonic/gin"
	"music_api_server/apiError"
	apiReq "music_api_server/request"
	service "music_api_server/service/user_service"
	"music_api_server/validator"
	"net/http"
)


//RegisterHandler TODO - USING CUSTOM ERROR HANDLER
func RegisterHandler(ctx *gin.Context) (interface{},error){
	req := apiReq.RegisterRequest{}

	//TODO -Binding the request
	err := ctx.ShouldBind(&req)
	if err != nil{
		return nil,apiError.APIError{
			Status: -1,
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	//TODO -Validator field Checking
	err = validator.Validate.Var(req.Email,"email")
	if err != nil {
			return nil,apiError.APIError{
					Status: http.StatusBadRequest,
					Code: http.StatusBadRequest,
					Message: "Please provide a email",
			}

	}
	//TODO - USer Service Register
	userService := service.UserService{}
	err = userService.Register(&req)
	if err != nil {
		return nil,apiError.APIError{
			Status: http.StatusBadRequest,
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return apiError.APIError{
		Status: 200,
		Message: "Register succeed!",
		Code: 200,
	}, nil
}

//LoginHandler TODO - USING CUSTOM ERROR HANDLER
func LoginHandler(ctx *gin.Context) (interface{},error){
	return nil, nil
}

//
//
//func tokenGenerator(ctx *gin.Context,info *user_service.UserInfo){
//	token := mid.NewJsonWebToken()
//	claims := mid.CustomClaims{
//		Name:  info.Name,
//		Email: info.Email,
//		StandardClaims: jwt.StandardClaims{
//			Issuer:"jackson.tmm",
//			ExpiresAt: int64(time.Now().Unix() + 3600),
//			NotBefore: int64(time.Now().Unix() - 1000),
//			Subject: "User Authorization",
//		},
//	}
//
//	jsonToken, err := token.CreateToken(claims)
//	if err != nil {
//		ctx.JSON(http.StatusOK,gin.H{
//			"status":-1,
//			"msg":err.Error(),
//			"data":nil,
//		})
//	}
//
//	log.Printf("User Token is %v",jsonToken)
//
//	ctx.JSON(http.StatusOK,gin.H{
//		"status":0,
//		"msg":"Logged in...",
//		"data": map[string]interface{}{
//			"user_service name":info.Name,
//			"token":jsonToken,
//		},
//	})
//	return
//}