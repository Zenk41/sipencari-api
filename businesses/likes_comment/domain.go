package likescomment

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	UserID    string
	UserName  string
	CommentID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetByID(idLike string, idMissing string, idComment int) Domain
	Like(idUser string, idMissing string,  idComment int, likeDomain *Domain) Domain
	Unlike(idUser string, idMissing string,  idComment int) bool
}

type Repository interface {
	GetAll() []Domain
	GetByID(idLike string, idMissing string, idComment int) Domain
	Like(idUser string, idMissing string,  idComment int, likeDomain *Domain) Domain
	Unlike(idUser string, idMissing string,  idComment int) bool
}
