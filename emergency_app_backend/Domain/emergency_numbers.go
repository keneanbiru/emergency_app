package Domain

import "context"

type EmergencyNumber struct {
	ID          string `json:"id"`          // Unique identifier
	Country     string `json:"country"`     // Country name
	Number      string `json:"number"`      // Emergency number
	Description string `json:"description"` // Description (e.g., police, ambulance)
}

type EmergencyNumberRepository interface {
	GetAllEmergencyNumbers(ctx context.Context) ([]EmergencyNumber, error)               // Fetch all emergency numbers
	SearchEmergencyNumbers(ctx context.Context, query string) ([]EmergencyNumber, error) // Search emergency numbers by country or description
}

type EmergencyNumberUsecaseInterface interface {
	GetEmergencyNumbers(ctx context.Context) ([]EmergencyNumber, error)
	SearchEmergencyNumbers(ctx context.Context, query string) ([]EmergencyNumber, error)
}
