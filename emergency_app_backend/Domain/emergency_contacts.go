package Domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmergencyContact represents an emergency contact
type EmergencyContact struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"` // Unique identifier (MongoDB ObjectId)
	UserID    string             `json:"user_id"`        // Associated user ID
	Name      string             `json:"name"`           // Contact name
	Phone     string             `json:"phone"`          // Phone number
	Relation  string             `json:"relation"`       // Relationship with the user
	CreatedAt string             `json:"created_at"`     // Timestamp of creation
	UpdatedAt string             `json:"updated_at"`     // Timestamp of the last update
}

// Convert ObjectId to string for JSON serialization
// func (ec *EmergencyContact) SetID(id primitive.ObjectID) {
// 	ec.ID = id
// }

// // Convert EmergencyContact to a JSON-friendly structure
// func (ec *EmergencyContact) ToJSON() map[string]interface{} {
// 	return map[string]interface{}{
// 		// "_id":        ec.ID, // Convert ObjectId to string for JSON serialization
// 		"user_id":    ec.UserID,
// 		"name":       ec.Name,
// 		"phone":      ec.Phone,
// 		"relation":   ec.Relation,
// 		"created_at": ec.CreatedAt,
// 		"updated_at": ec.UpdatedAt,
// 	}
// }

type EmergencyContactRepository interface {
	CreateEmergencyContact(ctx context.Context, contact EmergencyContact) (string, error) // Create a new contact
	GetEmergencyContactByID(ctx context.Context, id string) (*EmergencyContact, error)    // Retrieve contact by ID
	UpdateEmergencyContact(ctx context.Context, contact EmergencyContact) error           // Update an existing contact
	DeleteEmergencyContact(ctx context.Context, id string) error                          // Delete a contact by ID
	GetAllContacts(ctx context.Context) ([]EmergencyContact, error)                       // Get all contacts for a specific user
}

type EmergencyContactUsecaseInterface interface {
	CreateContact(ctx context.Context, contact EmergencyContact) (string, error)
	GetContactByID(ctx context.Context, id string) (*EmergencyContact, error)
	UpdateContact(ctx context.Context, contact EmergencyContact) error
	DeleteContact(ctx context.Context, id string) error
	GetUserContacts(ctx context.Context, userID string) ([]EmergencyContact, error)
}
