package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import  test "music_api_server/validator"

func TestRoute(r *gin.Engine){
	r.GET("/test", func(ctx *gin.Context) {
		var service test.TestService
		var user test.TestUser
		err := ctx.ShouldBind(&user)
		if err != nil{
			ctx.JSON(http.StatusOK,gin.H{
				"msg":test.Translate(err),
			})
			ctx.Abort()
			return
		}
		service = &user

		result := service.Greeting()
		ctx.JSON(http.StatusOK,gin.H{
			"msg":result,
		})
	})
}