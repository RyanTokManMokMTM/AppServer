package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTw "github.com/go-playground/validator/v10/translations/zh_tw"

	"music_api_server/route"
)

const (
	ADDR  string = "127.0.0.1"
	PORT int = 8080
)

var (
	universal *ut.UniversalTranslator
	validate *validator.Validate
	trans  ut.Translator
)


func main(){
	server := gin.Default()
	//TODO -Serving Static files
	resource := server.Group("/resource")
	resource.Static("/","./public")

	route.RouterInit(server) //init all available route
	log.Fatalln(server.Run(fmt.Sprintf("%s:%d",ADDR,PORT)))
}

//func ToolTest(){
//	fmt.Println(Tool.Base64KeysGenerator(64))
//}

func configTranslator(eng *gin.Engine){
	tw := zh_Hant_TW.New()
	universal = ut.New(tw,tw)

	trans , _ = universal.GetTranslator("zh_tw")

	validate = validator.New()
	zhTw.RegisterDefaultTranslations(validate,trans)
}

func Translate(err error) map[string][]string{
	result := make(map[string][]string)

	errs := err.(validator.ValidationErrors)
	for _,err:= range errs{
		result[err.Field()] = append(result[err.Field()],err.Translate(trans))
	}
	return  result
}


