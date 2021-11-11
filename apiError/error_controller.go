package apiError

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorController(ctx *gin.Context) (interface{},error){
	//test first
	query := ctx.Query("name")
	if query == ""{
		return nil,APIError{
			Status: http.StatusOK,
			Code: 400,
			Message: "please provide a name",
		}
	}

	if query == "admin"{
		return nil,APIError{
			Status: http.StatusOK,
			Code: 404,
			Message: "name is not allowed",
		}
	}

	return query,nil
}