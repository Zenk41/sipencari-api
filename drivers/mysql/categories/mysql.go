package categories

import (
	"sipencari-api/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll() []categories.Domain {
	var rec []Category

	cr.conn.Find(&rec)

	categoryDomain := []categories.Domain{}

	for _, category := range rec {
		categoryDomain = append(categoryDomain, category.ToDomain())
	}
	return categoryDomain
}

func (cr *categoryRepository) GetByID(category_id string) categories.Domain {
	var category Category

	cr.conn.First(&category, "id=?", category_id)

	return category.ToDomain()
}

func (cr *categoryRepository) Create(categoryDomain *categories.Domain) categories.Domain {
	rec := FromDomain(categoryDomain)

	result := cr.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (cr *categoryRepository) Update(category_id string, categoryDomain *categories.Domain) categories.Domain {
	var category categories.Domain = cr.GetByID(category_id)

	updatedCategory := FromDomain(&category)
	updatedCategory.Name = categoryDomain.Name
	updatedCategory.Description = categoryDomain.Description

	cr.conn.Save(&updatedCategory)

	return updatedCategory.ToDomain()
}

func (cr *categoryRepository) Delete(category_id string) bool {
	var category categories.Domain = cr.GetByID(category_id)

	deletedCategory := FromDomain(&category)
	if result := cr.conn.Unscoped().Delete(&deletedCategory); result.RowsAffected == 0 {
		return false
	}
	return true
}
