package locationsmissing

import (
	"sipencari-api/businesses/locations_missing"
	"time"

	"gorm.io/gorm"
)

type LocationMissing struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Lat       string         `json:"lat"`
	Lng       string         `json:"lng"`
	MissingID string         `json:"missing_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *locationsmissing.Domain) *LocationMissing {
	return &LocationMissing{
		ID:        domain.ID,
		Name:      domain.Name,
		Lat:       domain.Lat,
		Lng:       domain.Lng,
		MissingID: domain.MissingID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *LocationMissing) ToDomain() locationsmissing.Domain {
	return locationsmissing.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Lat:       rec.Lat,
		Lng:       rec.Lng,
		MissingID:rec.MissingID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
