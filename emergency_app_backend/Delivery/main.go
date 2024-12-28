package main

import (
	routers "emergency_app_backend/Delivery/routers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection setup
	clientOptions := options.Client().ApplyURI("mongodb+srv://kenean:Godislove33.@cluster0.fek5tj1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
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

	// Apply CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Change "*" to specific domains for better security
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router.Use(func(c *gin.Context) {
		corsHandler.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	// Set up routes
	routers.SetupRoutes(router, db)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no port is set
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
