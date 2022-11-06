package likescomment

import (
	likescomment "sipencari-api/businesses/likes_comment"
	"sipencari-api/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type LikeComment struct {
	User      users.User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID    string         `json:"user_id" gorm:"size:100" gorm:"primaryKey"`
	CommentID uint           `json:"comment_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *likescomment.Domain) *LikeComment {
	return &LikeComment{
		UserID:    domain.UserID,
		CommentID: domain.CommentID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *LikeComment) ToDomain() likescomment.Domain {
	return likescomment.Domain{
		UserID:    rec.UserID,
		UserName:  rec.User.Name,
		CommentID: rec.CommentID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
