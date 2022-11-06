package likesmissing

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	UserID    string
	UserName  string
	MissingID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll(idMissing string) []Domain
	GetByID(idLikes string, idMissing string) Domain
	Like(idUser string, idMissing string, likeDomain *Domain) Domain
	Unlike(idUser string, idMissing string) bool
}

type Repository interface {
	GetAll(idMissing string) []Domain
	GetByID(idLike string, idMissing string) Domain
	Like(idUser string, idMissing string, likeDomain *Domain) Domain
	Unlike(idUser string, idMissing string) bool
}
