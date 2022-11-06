package comments

import (
	likescomment "sipencari-api/businesses/likes_comment"
	locations "sipencari-api/businesses/locations_comment"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID              uint
	Message         string
	Image           string
	UserID          string
	UserName        string
	LocationComment locations.Domain
	LikeComments    []likescomment.Domain
	MissingID       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type Usecase interface {
	GetAll(idPost string) []Domain
	GetByID(idPost string, id string) Domain
	Create(idUser string, idPost string, commentDomain *Domain) Domain
	Update(idUser string, idPost string, idComment string, commentDomain *Domain) Domain
	Delete(idPost string, idComment string) bool
}

type Repository interface {
	GetAll(idPost string) []Domain
	GetByID(idPost string, id string) Domain
	Create(idUser string, idPost string, commentDomain *Domain) Domain
	Update(idUser string, idPost string, idComment string, commentDomain *Domain) Domain
	Delete(idPost string, idComment string) bool
}
