package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Picture   string
	IsAdmin   bool
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Register(userDomain *Domain) Domain
	CreateAdmin(userDomain *Domain) Domain
	Update(id string,userDomain *Domain) Domain
	Delete(id string) bool
	Login(userDomain *Domain) string
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Register(userDomain *Domain) Domain
	CreateAdmin(userDomain *Domain) Domain
	Update(id string,userDomain *Domain) Domain
	Delete(id string) bool
	Login(userDomain *Domain) Domain
}
