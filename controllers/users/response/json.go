package response

import (
	"sipencari-api/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Picture   string         `json:"picture"`
	IsAdmin   bool           `json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain users.Domain) User {
	return User{
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
