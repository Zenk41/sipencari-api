package locationsmissing

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	Name      string
	Lat       string
	Lng       string
	MissingID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetByID(idMissing, idLocation string) Domain
	Create(idMissing string, LocationMissingDomain *Domain) Domain
	Update(idMissing, idLocation string, LocationMissingDomain *Domain) Domain
	Delete(idMissing, idLocation string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(idMissing, idLocation string) Domain
	Create(idMissing string, LocationMissingDomain *Domain) Domain
	Update(idMissing, idLocation string, LocationMissingDomain *Domain) Domain
	Delete(idMissing, idLocation string) bool
}
