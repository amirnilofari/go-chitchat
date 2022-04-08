package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.SetTrustedProxies([]string{"192.168.1.2"})
	app.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "start with gin web framework",
		})

	})
	app.Run()
}
