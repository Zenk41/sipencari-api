package likesmissing

import (
	likesmissing "sipencari-api/businesses/likes_missing"
	"sipencari-api/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type LikeMissing struct {
	User      users.User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID    string         `json:"user_id" gorm:"size:100" gorm:"primaryKey"`
	MissingID string         `json:"missing_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *likesmissing.Domain) *LikeMissing {
	return &LikeMissing{
		UserID:    domain.UserID,
		MissingID: domain.MissingID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *LikeMissing) ToDomain() likesmissing.Domain {
	return likesmissing.Domain{
		UserID:    rec.UserID,
		UserName:  rec.User.Name,
		MissingID: rec.MissingID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
