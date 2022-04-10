package main

import (
	"net/http"

	"github.com/amirnilofari/go-chitchat/models"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	var posts []models.Post

	collection :=
		c.JSON(200, p1)
}

func err(response http.ResponseWriter, request *http.Request) {}

func login(response http.ResponseWriter, request *http.Request) {}

func logout(response http.ResponseWriter, request *http.Request) {}

func signup(response http.ResponseWriter, request *http.Request) {}

func signupAccount(response http.ResponseWriter, request *http.Request) {}

func authenticate(response http.ResponseWriter, request *http.Request) {}

func newThread(response http.ResponseWriter, request *http.Request) {}

func createThread(response http.ResponseWriter, request *http.Request) {}

func postThread(response http.ResponseWriter, request *http.Request) {}

func readThread(response http.ResponseWriter, request *http.Request) {}
