package routes

import (
	"github.com/amirnilofari/go-chitchat/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.POST("/v1/api/posts", controllers.CreatePost())
	router.GET("/v1/api/posts/:postId", controllers.GetAPost())
	router.PUT("/v1/api/posts/:postId", controllers.EditAPost())
	router.DELETE("/v1/api/posts/:postId", controllers.DeleteAUser())
	router.GET("/v1/api/posts", controllers.GetAllPost())
}
