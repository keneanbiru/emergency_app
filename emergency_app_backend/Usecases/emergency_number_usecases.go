package Usecases

import (
	"context"
	"emergency_app_backend/Domain"
)

// EmergencyNumberUsecase provides methods for managing emergency numbers.
type EmergencyNumberUsecase struct {
	repo Domain.EmergencyNumberRepository
}

// NewEmergencyNumberUsecase creates a new EmergencyNumberUsecase.
func NewEmergencyNumberUsecase(repo Domain.EmergencyNumberRepository) *EmergencyNumberUsecase {
	return &EmergencyNumberUsecase{repo: repo}
}

// GetEmergencyNumbers retrieves all emergency numbers.
func (u *EmergencyNumberUsecase) GetEmergencyNumbers(ctx context.Context) ([]Domain.EmergencyNumber, error) {
	return u.repo.GetAllEmergencyNumbers(ctx)
}

// SearchEmergencyNumbers searches emergency numbers based on a query.
func (u *EmergencyNumberUsecase) SearchEmergencyNumbers(ctx context.Context, query string) ([]Domain.EmergencyNumber, error) {
	return u.repo.SearchEmergencyNumbers(ctx, query)
}
