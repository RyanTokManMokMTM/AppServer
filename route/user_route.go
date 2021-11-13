package route

import (
	"github.com/gin-gonic/gin"
	"music_api_server/apiError"
	"music_api_server/controller/user_controller"
)

func UseRoute(r *gin.RouterGroup){
	userRoute := r.Group("/user/auth")
	userRoute.POST("/signup", apiError.ErrorHandler(user_controller.RegisterHandler))
	userRoute.POST("/login",apiError.ErrorHandler(user_controller.LoginHandler))
}
