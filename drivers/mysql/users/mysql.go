package users

import (
	"fmt"
	"sipencari-api/businesses/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) GetAll() []users.Domain {
	var rec []User

	ur.conn.Find(&rec)

	userDomain := []users.Domain{}

	for _, user := range rec {
		userDomain = append(userDomain, user.ToDomain())
	}
	return userDomain
}

func (ur *userRepository) GetByID(id string) users.Domain {
	var user User
	ur.conn.First(&user, "id=?", id)
	return user.ToDomain()
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
	id := uuid.New().String
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.ID = id()
	rec.Password = string(password)
	rec.IsAdmin = false
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) CreateAdmin(userDomain *users.Domain) users.Domain {
 id := uuid.New().String
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.ID = id()
	rec.Password = string(password)
	rec.IsAdmin = true
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) Update(id string,userDomain *users.Domain) users.Domain {
	var user users.Domain = ur.GetByID(id)
	updatedUser := FromDomain(&user)

	updatedUser.Name = userDomain.Name
	updatedUser.Email = userDomain.Email
	updatedUser.Password = userDomain.Password
	updatedUser.Picture = userDomain.Picture

	ur.conn.Save(&updatedUser)
	
	return updatedUser.ToDomain()
}

func (ur *userRepository) Delete(id string) bool {
	var user users.Domain = ur.GetByID(id)

	deletedUser := FromDomain(&user)

	if result := ur.conn.Delete(&deletedUser); result.RowsAffected == 0 {
		return false
	}
	return true
}

func (ur *userRepository) Login(userDomain *users.Domain) users.Domain {
	var user User
	ur.conn.First(&user, "email=?", userDomain.Email)

	if user.ID == "" {
		fmt.Println("user not found")
		return users.Domain{}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password)); err != nil {
		fmt.Println("wrong password")
		return users.Domain{}
	}
	return user.ToDomain()
}
