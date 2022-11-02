package mysql_driver

import (
	"fmt"
	"log"
	"sipencari-api/drivers/mysql/categories"
	"sipencari-api/drivers/mysql/users"

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
	db.AutoMigrate(&users.User{}, &categories.Category{})
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

// Seed
