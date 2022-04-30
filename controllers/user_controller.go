package controllers

import (
	"context"
	"github.com/amirnilofari/go-chitchat/Configs"
	"github.com/amirnilofari/go-chitchat/models"
	"github.com/amirnilofari/go-chitchat/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var postCollection *mongo.Collection = Configs.GetCollection(Configs.DB, "posts")
var validate = validator.New()

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var post models.Post
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&post); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest})
			return
		}

		newPost := models.Post{
			ID:          primitive.NewObjectID(),
			Title:       post.Title,
			Description: post.Description,
		}

		result, err := postCollection.InsertOne(ctx, newPost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		var post models.Post
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(postId)

		err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": post}})
	}
}

func EditAPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		var post models.Post
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(postId)

		//validate the request body
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		update := bson.M{"title": post.Title, "description": post.Description}
		result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//get updated post details
		var updatedPost models.Post
		if result.MatchedCount == 1 {
			err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedPost)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
		}
		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": updatedPost},
		})
	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := c.Param("postId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(postId)

		result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, responses.UserResponse{
				Status:  http.StatusNotFound,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": "Post successfully deleted!"},
		})
	}
}

func GetAllPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var posts []models.Post
		defer cancel()

		results, err := postCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
		}

		//reading from the DB in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singlePost models.Post
			if err = results.Decode(&singlePost); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})

			}
			posts = append(posts, singlePost)

		}
		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": posts},
		})
	}
}
