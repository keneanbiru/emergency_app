package Controllers

import (
	// "context"
	"emergency_app_backend/Domain"
	"emergency_app_backend/Usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmergencyContactController handles the HTTP requests related to emergency contacts.
type EmergencyContactController struct {
	usecase *Usecases.EmergencyContactUsecase
}

// NewEmergencyContactController creates a new EmergencyContactController.
// NewEmergencyContactController creates a new EmergencyContactController.
func NewEmergencyContactController(usecase *Usecases.EmergencyContactUsecase) *EmergencyContactController {
	return &EmergencyContactController{usecase: usecase}
}

// CreateContact handles the POST request to create a new emergency contact.
func (ctrl *EmergencyContactController) CreateContact(c *gin.Context) {
	var contact Domain.EmergencyContact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := ctrl.usecase.CreateContact(c, contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetContactByID handles the GET request to retrieve an emergency contact by ID.
func (ctrl *EmergencyContactController) GetContactByID(c *gin.Context) {
	id := c.Param("id")
	contact, err := ctrl.usecase.GetContactByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// UpdateContact handles the PUT request to update an emergency contact.
func (ctrl *EmergencyContactController) UpdateContact(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var contact Domain.EmergencyContact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact.ID = id
	if err := ctrl.usecase.UpdateContact(c, contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact updated successfully"})
}

// DeleteContact handles the DELETE request to delete an emergency contact.
func (ctrl *EmergencyContactController) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.usecase.DeleteContact(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}

// GetUserContacts handles the GET request to retrieve all emergency contacts for a user.
func (ctrl *EmergencyContactController) GetUserContacts(c *gin.Context) {
	// userID := c.Param("user_id")
	contacts, err := ctrl.usecase.GetUserContacts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}
