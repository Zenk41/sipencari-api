package hashtags

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Name      string
	MissingID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id_hashtag string) Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id_hashtag string) Domain
}
