package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"music_api_server/config"
	"music_api_server/route"
)

//ErrorHandler Return a gin.Handle function with error handle
func main() {

	server := gin.New()
	//TODO -Serving Static files
	resource := server.Group("/resource")
	resource.Static("/", "./public")

	route.RouterInit(server) //init all available route
	log.Fatalln(server.Run(fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)))
}
