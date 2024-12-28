package Usecases

import (
	"context"
	"emergency_app_backend/Domain"
)

// EmergencyContactUsecase provides methods for managing emergency contacts.
type EmergencyContactUsecase struct {
	repo Domain.EmergencyContactRepository
}

// NewEmergencyContactUsecase creates a new EmergencyContactUsecase.
func NewEmergencyContactUsecase(repo Domain.EmergencyContactRepository) *EmergencyContactUsecase {
	return &EmergencyContactUsecase{repo: repo}
}

// CreateContact adds a new emergency contact.
func (u *EmergencyContactUsecase) CreateContact(ctx context.Context, contact Domain.EmergencyContact) (string, error) {
	return u.repo.CreateEmergencyContact(ctx, contact)
}

// GetContactByID retrieves an emergency contact by ID.
func (u *EmergencyContactUsecase) GetContactByID(ctx context.Context, id string) (*Domain.EmergencyContact, error) {
	return u.repo.GetEmergencyContactByID(ctx, id)
}

// UpdateContact updates an existing emergency contact.
func (u *EmergencyContactUsecase) UpdateContact(ctx context.Context, contact Domain.EmergencyContact) error {
	return u.repo.UpdateEmergencyContact(ctx, contact)
}

// DeleteContact deletes an emergency contact by ID.
func (u *EmergencyContactUsecase) DeleteContact(ctx context.Context, id string) error {
	return u.repo.DeleteEmergencyContact(ctx, id)
}

// GetUserContacts retrieves all emergency contacts for a specific user.
func (u *EmergencyContactUsecase) GetUserContacts(ctx context.Context) ([]Domain.EmergencyContact, error) {
	return u.repo.GetAllContacts(ctx)
}
