package request

import (
	"sipencari-api/businesses/categories"

	"github.com/go-playground/validator"
)

type Category struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

func (req *Category) ToDomain() *categories.Domain {
	return &categories.Domain{
		Name: req.Name,
		Description: req.Description,
	}
}

func (req *Category) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	return err
}
