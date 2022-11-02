package drivers

import (
	userDomain "sipencari-api/businesses/users"
	userDB "sipencari-api/drivers/mysql/users"

	categoryDomain "sipencari-api/businesses/categories"
	categoryDB "sipencari-api/drivers/mysql/categories"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}