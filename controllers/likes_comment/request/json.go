package request

import (
	likescomment "sipencari-api/businesses/likes_comment"

	"github.com/go-playground/validator"
)

type LikeComment struct {
	UserID string `json:"user_id" form:"user_id" validate:"required"`
}

func (req *LikeComment) ToDomain() *likescomment.Domain {
	return &likescomment.Domain{
		UserID: req.UserID,
	}
}

func (req *LikeComment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}