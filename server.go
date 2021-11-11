package main

import (
	"fmt"
	"log"

	"music_api_server/config"
	"github.com/gin-gonic/gin"
	"music_api_server/route"
)


func main(){
	server := gin.Default()
	//TODO -Serving Static files
	resource := server.Group("/resource")
	resource.Static("/","./public")

	route.RouterInit(server) //init all available route

	log.Fatalln(server.Run(fmt.Sprintf("%s:%d",config.Server.Address,config.Server.Port)))
}


