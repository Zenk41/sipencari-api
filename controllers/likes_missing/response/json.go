package response

import (
	likesmissing "sipencari-api/businesses/likes_missing"
	"time"

	"gorm.io/gorm"
)

type LikeMissing struct {
	UserID    string         `json:"user_id" gorm:"primaryKey"`
	UserName  string         `json:"user_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain likesmissing.Domain) LikeMissing {
	return LikeMissing{
		UserID:    domain.UserID,
		UserName:  domain.UserName,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
