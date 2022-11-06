package missings

import (
	"sipencari-api/businesses/comments"
	"sipencari-api/businesses/hashtags"
	likesmissing "sipencari-api/businesses/likes_missing"
	locationsmissing "sipencari-api/businesses/locations_missing"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID              string
	Title           string
	Content         string
	Image           string
	UserID          string
	UserName        string
	LocationMissing locationsmissing.Domain
	CategoryID      uint
	CategoryName    string
	Hashtags        []hashtags.Domain
	LikeMissings    []likesmissing.Domain
	Comments        []comments.Domain
	IsFound         bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(idUser string, missingDomain *Domain) Domain
	Update(idUser string, id string, missingDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(idUser string, missingDomain *Domain) Domain
	Update(idUser string, id string, missingDomain *Domain) Domain
	Delete(id string) bool
}
