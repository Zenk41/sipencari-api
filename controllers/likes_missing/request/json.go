package request

import (
	likesmissing "sipencari-api/businesses/likes_missing"

	"github.com/go-playground/validator"
)

type LikeMissing struct {
	UserID string `json:"user_id" form:"user_id" validate:"required"`
}

func (req *LikeMissing) ToDomain() *likesmissing.Domain {
	return &likesmissing.Domain{
		UserID: req.UserID,
	}
}

func (req *LikeMissing) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}
