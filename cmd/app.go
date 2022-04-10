package main

import (
	"context"
	"time"

	"github.com/amirnilofari/go-chitchat/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI = "mongodb://localhost:27017"

func main() {
	app := gin.Default()
	app.SetTrustedProxies([]string{"192.168.1.2"})

	//Connect DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	app.GET("v1/api/posts", controllers.AllPost)
	app.GET("v1/api/posts", controllers.CreatePost)
	app.GET("v1/api/posts", controllers.DeletePost)

	app.Run()
}
