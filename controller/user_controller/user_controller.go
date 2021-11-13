package user_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music_api_server/apiError"
	apiReq "music_api_server/request"
	service "music_api_server/service/user_service"
	"music_api_server/validator"
	"net/http"
	"strings"
)


//RegisterHandler TODO - USING CUSTOM ERROR HANDLER
func RegisterHandler(ctx *gin.Context) (interface{},error){
	req := apiReq.RegisterRequest{}
	//TODO -Binding the request
	err := ctx.ShouldBind(&req)
	if err != nil{
		return nil,apiError.APIError{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if match := strings.Compare(req.Password,req.ConfirmPassword);match != 0{
		return nil ,apiError.APIError{
			Code:http.StatusOK,
			Message: "password not match",
		}
	}

	err = validator.Validate.Var(req.Password,"alphanum")
	if err != nil{
		fmt.Println(err)
		return nil,apiError.APIError{
			Code: http.StatusOK,
			Message: "Password must only contain alphabet and number",
		}
	}

	//TODO -Validator field Checking
	err = validator.Validate.Var(req.Email,"email")
	if err != nil {
			return nil,apiError.APIError{
					Code: http.StatusBadRequest,
					Message: "Please provide a email",
			}

	}
	//TODO - USer Service Register
	userService := service.UserService{}
	err = userService.Register(&req)
	if err != nil {
		return nil,apiError.APIError{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return "Register succeed!", nil
}

//LoginHandler TODO - USING CUSTOM ERROR HANDLER
func LoginHandler(ctx *gin.Context) (interface{},error){
	req := apiReq.LoginRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		return nil ,apiError.APIError{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	service := service.UserService{}
	jwt, err := service.Login(&req)
	if err != nil {
		return nil, apiError.APIError{
			Code: http.StatusOK,
			Message: err.Error(),
		}
	}
	authToken := fmt.Sprintf("Bearer %s",jwt)
	return authToken, nil
}
