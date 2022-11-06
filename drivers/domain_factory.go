package drivers

import (
	userDomain "sipencari-api/businesses/users"
	userDB "sipencari-api/drivers/mysql/users"

	categoryDomain "sipencari-api/businesses/categories"
	categoryDB "sipencari-api/drivers/mysql/categories"

	missingDomain "sipencari-api/businesses/missings"
	missingDB "sipencari-api/drivers/mysql/missings"

	commentDomain "sipencari-api/businesses/comments"
	commentDB "sipencari-api/drivers/mysql/comments"

	likeMissingDomain "sipencari-api/businesses/likes_missing"
	likeMissingDB "sipencari-api/drivers/mysql/likes_missing"

	likeCommentDomain "sipencari-api/businesses/likes_comment"
	likeCommentDB "sipencari-api/drivers/mysql/likes_comment"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewMissingRepository(conn *gorm.DB) missingDomain.Repository {
	return missingDB.NewMySQLRepository(conn)
}

func NewCommentRepository(conn *gorm.DB) commentDomain.Repository {
	return commentDB.NewMySQLRepository(conn)
}

func NewLikeMissingRepository(conn *gorm.DB) likeMissingDomain.Repository {
	return likeMissingDB.NewMySQLRepository(conn)
}

func NewLikeCommentRepository(conn *gorm.DB) likeCommentDomain.Repository {
	return likeCommentDB.NewMySQLRepository(conn)
}
