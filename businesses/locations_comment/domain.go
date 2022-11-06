package locationscomment

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	Name      string
	Lat       string
	Lng       string
	CommentID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	Create(idComment string, LocationMissingDomain *Domain) Domain
	GetByID(idComment, idLocation string) Domain
	Update(idComment, idLocation string, missingDomain *Domain) Domain
	Delete(idComment, idLocation string) bool
}

type Repository interface {
	GetAll() []Domain
	Create(idComment string, LocationMissingDomain *Domain) Domain
	GetByID(idComment, idLocation string) Domain
	Update(idComment, idLocation string, missingDomain *Domain) Domain
	Delete(idComment, idLocation string) bool
}