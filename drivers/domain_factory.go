package drivers

import (
	userDomain "sipencari-api/businesses/users"
	userDB "sipencari-api/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}