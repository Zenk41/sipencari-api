package mysql_driver

import (
	"errors"
	"fmt"
	"log"
	"sipencari-api/drivers/mysql/categories"
	"sipencari-api/drivers/mysql/comments"
	"sipencari-api/drivers/mysql/hashtags"
	likescomment "sipencari-api/drivers/mysql/likes_comment"
	likesmissing "sipencari-api/drivers/mysql/likes_missing"
	locationscomment "sipencari-api/drivers/mysql/locations_comment"
	locationsmissing "sipencari-api/drivers/mysql/locations_missing"
	"sipencari-api/drivers/mysql/missings"
	"sipencari-api/drivers/mysql/users"
	"sipencari-api/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database Configuration
type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

// Connecting to a Database server
func (conf *ConfigDB) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connecting to a database server : %s ", err)
	}
	log.Println("connected to a database server")
	return db
}

// Migrating struct into table
func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&users.User{},
		&categories.Category{},
		&missings.Missing{},
		&comments.Comment{},
		&locationscomment.LocationComment{},
		&locationsmissing.LocationMissing{},
		&hashtags.Hashtag{},
		&likescomment.LikeComment{},
		&likesmissing.LikeMissing{},
	)
}

// Closing Database
func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		log.Printf("error when getting the database instance : %v", err)
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection : %v", err)
	}
	log.Println("database connection is closed")
	return nil
}

// Seed for testing

func SeedUser(db *gorm.DB) users.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("testing"), bcrypt.DefaultCost)
	fakeUser, _ := utils.CreateFaker[users.User]()
	idUser := uuid.NewString()

	userRecord := users.User{
		ID: idUser,
		Name:     fakeUser.Name,
		Email:    fakeUser.Email,
		Password: string(password),
	}
	if err := db.Create(&userRecord).Error; err != nil {
		panic(err)
	}
	var lastUser users.User
	db.Last(&lastUser)

	lastUser.Password = "testing"
	return lastUser
}

func CleanSeeds(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	userResult := db.Exec("DELETE FROM users")

	if userResult.Error != nil {
		panic(errors.New("error when cleaning up users seeders"))
	}
	log.Println("Seeders are cleaned up successfully")
}
