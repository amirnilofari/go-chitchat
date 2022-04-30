package main

import (
	"github.com/amirnilofari/go-chitchat/Configs"
	"github.com/amirnilofari/go-chitchat/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	Configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	router.Run()

}
