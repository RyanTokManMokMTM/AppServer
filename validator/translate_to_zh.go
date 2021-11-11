package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTw "github.com/go-playground/validator/v10/translations/zh_tw"
)

var (
	universal *ut.UniversalTranslator
	trans  ut.Translator
)

type TestService interface {
	Greeting() string
}

type TestStr struct {
	Name string `form:"username" binding:"required"`
}

type TestUser struct {
	Name string `form:"username" binding:"required"`
}


func (test* TestStr) Greeting(){

}

func (u *TestUser) Greeting() string{
	return fmt.Sprintf("hello,%v,welcome!!!",u.Name)
}

func init(){
	tw := zh_Hant_TW.New() //new zh_tw translation
	universal = ut.New(tw,tw) //new translator

	trans , _ = universal.GetTranslator("zh_tw")

	//setting gin binding validator
	validate := binding.Validator.Engine().(*validator.Validate)

	//register translator to gin validator
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

