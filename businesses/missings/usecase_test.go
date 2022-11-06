package missings_test

import (
	// "sipencari-api/businesses/hashtags"
	// locations_missing "sipencari-api/businesses/locations_missing"
	"sipencari-api/businesses/missings"
	_missingMock "sipencari-api/businesses/missings/mocks"
	"sipencari-api/businesses/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	missingsRepository _missingMock.Repository
	missingsService    missings.Usecase
	usersDomain users.Domain
	missingsDomain missings.Domain
)

func TestMain(m *testing.M) {
	missingsService = missings.NewMissingUsecase(&missingsRepository)
	usersDomain = users.Domain{
		ID:       "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613",
		Name:     "testing",
		Email:    "testing123@testing.com",
		Password: "testing123",
	}
	missingsDomain = missings.Domain{
		ID:         "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613",
		Title:      "Anak Hilang 5 Tahun Setelah Pulang Sekolah",
		Content:    "disini gua mau coba cari adik gua yang udah hilang beberapa tahun. terakhir ada dirumah itu sekitar tahun 2015/2016. kita udah coba cari kemana2, bahkan udah lapor polisi juga, tapi tetep g ada hasil nya.",
		CategoryID: 1,
		UserID:     usersDomain.ID,
	}
	m.Run()
}


func TestGetAll(t *testing.T) {
	t.Run("TestGetAll | Valid", func(t *testing.T) {

		missingsRepository.On("GetAll").Return([]missings.Domain{missingsDomain}).Once()
		result := missingsService.GetAll()
		assert.NotNil(t, result)
	})
	t.Run("TestGetAll | InValid", func(t *testing.T) {
		missingsRepository.On("GetAll").Return([]missings.Domain{}).Once()
		result := missingsService.GetAll()
		assert.NotNil(t, 0, len(result))
	})
}


func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		missingsRepository.On("Create", usersDomain.ID, &missingsDomain).Return(missings.Domain{}).Once()
		result := missingsService.Create(usersDomain.ID, &missingsDomain)
		assert.NotNil(t, result)
	})
	t.Run("Register | InValid", func(t *testing.T) {
		missingsRepository.On("Create", "",&missings.Domain{}).Return(missings.Domain{}).Once()
		result := missingsService.Create("", &missings.Domain{})
		assert.NotNil(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {

		missingsRepository.On("GetByID", missingsDomain.ID).Return(missingsDomain).Once()
		result := missingsService.GetByID(missingsDomain.ID)
		assert.NotNil(t, result)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		missingsRepository.On("GetByID", "").Return(missings.Domain{}).Once()
		result := missingsService.GetByID("")
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		missingsRepository.On("Update", usersDomain.ID, missingsDomain.ID, &missingsDomain).Return(missings.Domain{}).Once()
		result := missingsService.Update(usersDomain.ID, missingsDomain.ID, &missingsDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		missingsRepository.On("Update", usersDomain.ID, missingsDomain.ID, &missings.Domain{}).Return(missings.Domain{}).Once()

		result := missingsService.Update(usersDomain.ID, missingsDomain.ID, &missings.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		missingsRepository.On("Delete", missingsDomain.ID).Return(true).Once()
		result := missingsService.Delete(missingsDomain.ID)

		assert.NotNil(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		missingsRepository.On("Delete", "").Return(false).Once()

		result := missingsService.Delete("")

		assert.NotNil(t, result)
	})
}
