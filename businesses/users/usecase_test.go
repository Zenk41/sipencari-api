package users_test

import (
	"sipencari-api/app/middlewares"
	"sipencari-api/businesses/users"
	_userMock "sipencari-api/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase
	usersDomain     users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepository, &middlewares.ConfigJWT{})
	usersDomain = users.Domain{
		ID:       "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613",
		Name:     "testing",
		Email:    "testing123@testing.com",
		Password: "testing123",
	}
	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		usersRepository.On("Register", &usersDomain).Return(users.Domain{}).Once()
		result := usersService.Register(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Register | InValid", func(t *testing.T) {
		usersRepository.On("Register", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Register(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestCreateAdmin(t *testing.T) {
	t.Run("CreateAdmin | Valid", func(t *testing.T) {
		usersRepository.On("CreateAdmin", &usersDomain).Return(users.Domain{}).Once()
		result := usersService.CreateAdmin(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("CreateAdmin | InValid", func(t *testing.T) {
		usersRepository.On("CreateAdmin", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.CreateAdmin(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		usersRepository.On("GetAll").Return([]users.Domain{usersDomain}).Once()
		result := usersService.GetAll()

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("Login", &usersDomain).Return(users.Domain{}).Once()

		result := usersService.Login(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("Login", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Login(&users.Domain{})

		assert.Empty(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		usersRepository.On("Delete", "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613").Return(true).Once()
		result := usersService.Delete("7a3a1ed7-b1c0-4079-a41a-f7253ddc2613")

		assert.NotNil(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		usersRepository.On("Delete", "").Return(false).Once()

		result := usersService.Delete("")

		assert.NotNil(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {

		usersRepository.On("GetByID", "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613").Return(usersDomain).Once()
		result := usersService.GetByID("7a3a1ed7-b1c0-4079-a41a-f7253ddc2613")
		assert.NotNil(t, result)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		usersRepository.On("GetByID", "").Return(users.Domain{}).Once()
		result := usersService.GetByID("")
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		idString := "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613"
		usersRepository.On("Update", idString, &usersDomain).Return(users.Domain{}).Once()
		result := usersService.Update(idString, &usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		usersRepository.On("Update", "", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Update("", &users.Domain{})

		assert.NotNil(t, result)
	})
}
