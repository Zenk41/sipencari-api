package response

import (
	locations "sipencari-api/businesses/locations_missing"
	"time"

	"gorm.io/gorm"
)

type LocationMissing struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Lat       string         `json:"lat"`
	Lng       string         `json:"lng"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain locations.Domain) LocationMissing {
	return LocationMissing{
		ID:        domain.ID,
		Name:      domain.Name,
		Lat:       domain.Lat,
		Lng:       domain.Lng,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
