package Controllers

import (
	"emergency_app_backend/Usecases"
	"net/http"

	// "emergency_app_backend/Domain"
	"github.com/gin-gonic/gin"
)

// EmergencyNumberController handles the HTTP requests related to emergency numbers.
type EmergencyNumberController struct {
	usecase *Usecases.EmergencyNumberUsecase
}

// NewEmergencyNumberController creates a new EmergencyNumberController.
// NewEmergencyNumberController creates a new EmergencyNumberController.
func NewEmergencyNumberController(usecase *Usecases.EmergencyNumberUsecase) *EmergencyNumberController {
	return &EmergencyNumberController{usecase: usecase}
}

// GetEmergencyNumbers handles the GET request to retrieve all emergency numbers.
func (ctrl *EmergencyNumberController) GetEmergencyNumbers(c *gin.Context) {
	numbers, err := ctrl.usecase.GetEmergencyNumbers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, numbers)
}

// SearchEmergencyNumbers handles the GET request to search emergency numbers based on a query.
func (ctrl *EmergencyNumberController) SearchEmergencyNumbers(c *gin.Context) {
	query := c.DefaultQuery("query", "")
	numbers, err := ctrl.usecase.SearchEmergencyNumbers(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, numbers)
}
