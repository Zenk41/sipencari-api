package categories

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetByID(category_id string) Domain
	Create(categoryDomain *Domain) Domain
	Update(category_id string, categoryDomain *Domain) Domain
	Delete(category_id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(category_id string) Domain
	Create(categoryDomain *Domain) Domain
	Update(category_id string, categoryDomain *Domain) Domain
	Delete(category_id string) bool
}
