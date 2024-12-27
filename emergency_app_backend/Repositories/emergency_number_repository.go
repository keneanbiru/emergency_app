package Repositories

import (
	"context"
	"emergency_app_backend/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmergencyNumberRepo struct {
	collection *mongo.Collection
}

// Constructor for EmergencyNumberRepo
func NewEmergencyNumberRepo(collection *mongo.Collection) *EmergencyNumberRepo {
	return &EmergencyNumberRepo{collection: collection}
}

// Ensure EmergencyNumberRepo implements Domain.EmergencyNumberRepository
var _ Domain.EmergencyNumberRepository = &EmergencyNumberRepo{}

func (repo *EmergencyNumberRepo) GetAllEmergencyNumbers(ctx context.Context) ([]Domain.EmergencyNumber, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var numbers []Domain.EmergencyNumber
	if err := cursor.All(ctx, &numbers); err != nil {
		return nil, err
	}
	return numbers, nil
}

func (repo *EmergencyNumberRepo) SearchEmergencyNumbers(ctx context.Context, query string) ([]Domain.EmergencyNumber, error) {
	filter := bson.M{"$or": []bson.M{
		{"country": bson.M{"$regex": query, "$options": "i"}},
		{"description": bson.M{"$regex": query, "$options": "i"}},
	}}
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var numbers []Domain.EmergencyNumber
	if err := cursor.All(ctx, &numbers); err != nil {
		return nil, err
	}
	return numbers, nil
}
