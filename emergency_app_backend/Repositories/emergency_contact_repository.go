package Repositories

import (
	"context"
	"emergency_app_backend/Domain"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmergencyContactRepo struct {
	collection *mongo.Collection
}

// Constructor for EmergencyContactRepo
func NewEmergencyContactRepo(collection *mongo.Collection) *EmergencyContactRepo {
	return &EmergencyContactRepo{collection: collection}
}

// Ensure EmergencyContactRepo implements Domain.EmergencyContactRepository
var _ Domain.EmergencyContactRepository = &EmergencyContactRepo{}

func (repo *EmergencyContactRepo) CreateEmergencyContact(ctx context.Context, contact Domain.EmergencyContact) (string, error) {
	result, err := repo.collection.InsertOne(ctx, contact)
	if err != nil {
		return "", err
	}

	// Safely convert InsertedID to a string
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to cast InsertedID to ObjectID")
	}

	return objectID.Hex(), nil
}

func (repo *EmergencyContactRepo) GetEmergencyContactByID(ctx context.Context, id string) (*Domain.EmergencyContact, error) {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Query the database
	var contact Domain.EmergencyContact
	err = repo.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&contact)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("contact not found")
		}
		return nil, err
	}

	return &contact, nil
}

func (repo *EmergencyContactRepo) UpdateEmergencyContact(ctx context.Context, contact Domain.EmergencyContact) error {
	// Convert the contact ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(contact.ID)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	// Create the filter and update query
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": contact}

	// Perform the update
	result, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Check if any document was updated
	if result.MatchedCount == 0 {
		return fmt.Errorf("contact not found")
	}

	return nil
}

func (repo *EmergencyContactRepo) DeleteEmergencyContact(ctx context.Context, id string) error {
	// Convert the string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	// Perform the delete operation
	result, err := repo.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("contact not found")
	}

	return nil
}

func (repo *EmergencyContactRepo) GetContactsByUserID(ctx context.Context, userID string) ([]Domain.EmergencyContact, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	var contacts []Domain.EmergencyContact
	if err := cursor.All(ctx, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}