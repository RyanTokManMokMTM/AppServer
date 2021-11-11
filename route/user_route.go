package route

import (
	"github.com/gin-gonic/gin"
	ctr "music_api_server/controller"
)

func UseRoute(r *gin.RouterGroup){
	userRoute := r.Group("/user/auth")
	userRoute.POST("/login",ctr.UserLogin)
	userRoute.POST("/signup",ctr.UserSignUp)
	userRoute.POST("/logout",ctr.UserSignOut)
}
