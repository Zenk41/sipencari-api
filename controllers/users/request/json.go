package request

import (
	"sipencari-api/businesses/users"

	"github.com/go-playground/validator"
)

type User struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Picture  string `json:"picture" form:"picture"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Picture:  req.Picture,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}


func (req *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)
	return err
}