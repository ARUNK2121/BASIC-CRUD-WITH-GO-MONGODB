package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-mongo-sample/delivery"
	"go-mongo-sample/repository"
)

func main() {
	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repository := repository.NewRepository(client.Database("users"))

	// create an http server
	handler := delivery.NewHandler(repository)

	// create a gin router
	router := gin.Default()
	{
		router.GET("/users/:email", handler.GetUser)
		router.POST("/users", handler.CreateUser)
		router.PUT("/users/:email", handler.UpdateUser)
		router.DELETE("/users/:email", handler.DeleteUser)
	}

	// start the router
	router.Run(":3000")
}
