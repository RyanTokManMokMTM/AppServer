package route

import (
	"github.com/gin-gonic/gin"
)

// RouterInit - /api/v1/path....
func RouterInit(r *gin.Engine){
	//Init for all available route
	apiV1 := r.Group("/api/v1")
	UseRoute(apiV1) //User route
	//r.GET("/testErr",apiError.ErrHandler(apiError.ErrController))
	//TestRoute(r)

}

//func logger() gin.HandlerFunc{
//	return func(ctx* gin.Context){
//		t := time.Time{}
//		ctx.Set("test_field","jackson testing")
//
//		ctx.Next()
//
//		latency := time.Since(t)
//		fmt.Println(latency)
//		log.Println(ctx.Writer.Status())
//	}
//}