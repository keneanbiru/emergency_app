package main

import (
	routers "emergency_app_backend/Delivery/routers"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "net/http"
)

func main() {
	// MongoDB connection setup
	clientOptions := options.Client().ApplyURI("mongodb+srv://<><>@cluster0.fek5tj1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check MongoDB connection
	err = client.Ping(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get the database instance
	db := client.Database("emergencyApp")

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routers.SetupRoutes(router, db)

	// Start the server
	router.Run(":8080")
}
