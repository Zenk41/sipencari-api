package users

import (
	"sipencari-api/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"size:100;primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Picture   string         `json:"picture"`
	IsAdmin   bool           `json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Picture:   domain.Picture,
		IsAdmin:   domain.IsAdmin,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdateAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Picture:   rec.Picture,
		IsAdmin:   rec.IsAdmin,
		CreatedAt: rec.CreatedAt,
		UpdateAt:  rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
