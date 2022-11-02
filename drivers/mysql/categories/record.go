package categories

import (
	"sipencari-api/businesses/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (rec *Category) ToDomain() categories.Domain {
	return categories.Domain{
		ID:          rec.ID,
		Name:        rec.Name,
		Description: rec.Description,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
	}
}

func FromDomain(domain *categories.Domain) *Category {
	return &Category{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}
