package locationscomment

import (
	"sipencari-api/businesses/locations_comment"
	"time"

	"gorm.io/gorm"
)

type LocationComment struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Lat       string         `json:"lat"`
	Lng       string         `json:"lng"`
	CommentID uint           `json:"comment_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *locationscomment.Domain) *LocationComment {
	return &LocationComment{
		ID:        domain.ID,
		Name:      domain.Name,
		Lat:       domain.Lat,
		Lng:       domain.Lng,
		CommentID: domain.CommentID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *LocationComment) ToDomain() locationscomment.Domain {
	return locationscomment.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Lat:       rec.Lat,
		Lng:       rec.Lng,
		CommentID: rec.CommentID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
