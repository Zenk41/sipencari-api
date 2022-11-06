package response

import (
	"sipencari-api/businesses/comments"
	likescomment "sipencari-api/businesses/likes_comment"
	resLikeComment "sipencari-api/controllers/likes_comment/response"
	resLocation "sipencari-api/controllers/locations_comment/response"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID              uint                         `json:"id" gorm:"primaryKey"`
	Message         string                       `json:"message"`
	Image           string                       `json:"image"`
	UserID          string                       `json:"user_id"`
	UserName        string                       `json:"user_name"`
	LocationComment resLocation.LocationComment  `json:"location_comment"`
	LikeComments    []resLikeComment.LikeComment `json:"like_comments"`
	CreatedAt       time.Time                    `json:"created_at"`
	UpdatedAt       time.Time                    `json:"updated_at"`
	DeletedAt       gorm.DeletedAt               `json:"deleted_at"`
}

func FromDomain(domain comments.Domain) Comment {
	var likesCommentData []resLikeComment.LikeComment
	likescommentToDomain := domain.LikeComments
	for _, likecomment := range likescommentToDomain {
		likesCommentData = append(likesCommentData, resLikeComment.FromDomain(likescomment.Domain{
			UserID:   likecomment.UserID,
			UserName: likecomment.UserName,
		}))
	}

	return Comment{
		ID:       domain.ID,
		Message:  domain.Message,
		Image:    domain.Image,
		UserID:   domain.UserID,
		UserName: domain.UserName,
		LocationComment: resLocation.LocationComment{
			ID:   domain.LocationComment.ID,
			Name: domain.LocationComment.Name,
			Lat:  domain.LocationComment.Lat,
			Lng:  domain.LocationComment.Lng,
		},
		LikeComments: likesCommentData,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
	}
}
