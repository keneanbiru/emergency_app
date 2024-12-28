package Router

import (
	"emergency_app_backend/Delivery/Controllers"
	"emergency_app_backend/Repositories"
	"emergency_app_backend/Usecases"

	// "emergency_app_backend/Domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(router *gin.Engine, db *mongo.Database) {
	// Repositories
	emergencyContactRepo := Repositories.NewEmergencyContactRepo(db.Collection("emergencyContacts"))
	emergencyNumberRepo := Repositories.NewEmergencyNumberRepo(db.Collection("emergencyNumbers"))

	// Usecases
	emergencyContactUsecase := Usecases.NewEmergencyContactUsecase(emergencyContactRepo)
	emergencyNumberUsecase := Usecases.NewEmergencyNumberUsecase(emergencyNumberRepo)

	// Controllers
	emergencyContactController := Controllers.NewEmergencyContactController(emergencyContactUsecase)
	emergencyNumberController := Controllers.NewEmergencyNumberController(emergencyNumberUsecase)

	// Emergency Contact Routes
	emergencyContactRoutes := router.Group("/emergency-contacts")
	{
		emergencyContactRoutes.POST("/", emergencyContactController.CreateContact)
		emergencyContactRoutes.GET("/:id", emergencyContactController.GetContactByID)
		emergencyContactRoutes.PUT("/:id", emergencyContactController.UpdateContact)
		emergencyContactRoutes.DELETE("/:id", emergencyContactController.DeleteContact)
		emergencyContactRoutes.GET("/", emergencyContactController.GetUserContacts)
	}

	// Emergency Number Routes``
	emergencyNumberRoutes := router.Group("/emergency-numbers")
	{
		emergencyNumberRoutes.GET("/", emergencyNumberController.GetEmergencyNumbers)
		emergencyNumberRoutes.GET("/search", emergencyNumberController.SearchEmergencyNumbers)
	}
}
