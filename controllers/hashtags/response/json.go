package response

import (
	"sipencari-api/businesses/hashtags"
	"time"

	"gorm.io/gorm"
)

type Hashtag struct {
	Name      string         `json:"name" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain hashtags.Domain) Hashtag {
	return Hashtag{
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
