package route

import (
	"github.com/gin-gonic/gin"
)

// RouterInit - /api/v1/path....
func RouterInit(r *gin.Engine){
	//Init for all available route
	apiV1 := r.Group("/api/v1")
	UseRoute(apiV1) //User route
	TestRoute(r)
}
