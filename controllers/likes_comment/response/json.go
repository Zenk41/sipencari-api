package response

import (
	likescomment "sipencari-api/businesses/likes_comment"
	"time"

	"gorm.io/gorm"
)

type LikeComment struct {
	UserID    string         `json:"user_id" gorm:"primaryKey"`
	UserName  string         `json:"user_name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain likescomment.Domain) LikeComment {
	return LikeComment{
		UserID:    domain.UserID,
		UserName:  domain.UserName,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
