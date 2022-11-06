package comments

import (
	"sipencari-api/businesses/comments"
	mysqllikes "sipencari-api/drivers/mysql/likes_comment"
	mysqllocations "sipencari-api/drivers/mysql/locations_comment"
	busLikes "sipencari-api/businesses/likes_comment"
	"sipencari-api/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID              uint                           `json:"id" gorm:"primaryKey"`
	Message         string                         `json:"message"`
	Image           string                         `json:"image"`
	UserID          string                         `json:"user_id" gorm:"size:100"`
	User            users.User                     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LocationComment mysqllocations.LocationComment `json:"location_comment" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LikeComments    []mysqllikes.LikeComment       `json:"like_comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MissingID       string                         `json:"missing_id"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt                 `json:"deleted_at"`
}

func FromDomain(domain *comments.Domain) *Comment {
	var likesData []mysqllikes.LikeComment
	likesFromDomain := domain.LikeComments
	for _, like := range likesFromDomain {
		likesData = append(likesData, mysqllikes.LikeComment{
			UserID: like.UserID,
		})
	}
	return &Comment{
		ID:      domain.ID,
		Message: domain.Message,
		Image:   domain.Image,
		UserID:  domain.UserID,
		LocationComment: mysqllocations.LocationComment{
			ID:   domain.LocationComment.ID,
			Name: domain.LocationComment.Name,
			Lat:  domain.LocationComment.Lat,
			Lng:  domain.LocationComment.Lng,
		},
		LikeComments: likesData,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}

func (rec *Comment) ToDomain() comments.Domain {

	var likesFromDomain []busLikes.Domain
	for _, like := range rec.LikeComments {
		likesFromDomain = append(likesFromDomain, like.ToDomain())
	}
	return comments.Domain{
		ID:              rec.ID,
		Message:         rec.Message,
		Image:           rec.Image,
		UserID:          rec.UserID,
		UserName:        rec.User.Name,
		LocationComment: rec.LocationComment.ToDomain(),
		LikeComments:    likesFromDomain,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
		DeletedAt:       rec.DeletedAt,
	}
}
