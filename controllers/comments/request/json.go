package request

import (
	"sipencari-api/businesses/comments"
	likescomment "sipencari-api/businesses/likes_comment"
	locationscomment "sipencari-api/businesses/locations_comment"
	reqLikeComment "sipencari-api/controllers/likes_comment/request"
	reqLocationComment "sipencari-api/controllers/locations_comment/request"

	"github.com/go-playground/validator"
)

type Comment struct {
	Message         string                             `json:"message" form:"message" validate:"required"`
	Image           string                             `json:"image" form:"image"`
	UserID          string                             `json:"user_id" form:"user_id"`
	LocationComment reqLocationComment.LocationComment `json:"location_comment" form:"location_comment" validate:"required"`
	LikeComments    []reqLikeComment.LikeComment       `json:"like_comments" form:"like_comments"`
}

func (req *Comment) ToDomain() *comments.Domain {

	var likescomment []likescomment.Domain
	for _, likecomment := range req.LikeComments {
		likescomment = append(likescomment, *likecomment.ToDomain())
	}

	return &comments.Domain{
		Message: req.Message,
		Image:   req.Image,
		UserID:  req.UserID,
		LocationComment: locationscomment.Domain{
			Name: req.LocationComment.Name,
			Lat:  req.LocationComment.Lat,
			Lng:  req.LocationComment.Lng,
		},
		LikeComments: likescomment,
	}
}

func (req *Comment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}