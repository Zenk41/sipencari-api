package likesmissing_test

import (

	likesmissing "sipencari-api/businesses/likes_missing"
	_likeMissingMock "sipencari-api/businesses/likes_missing/mocks"
	"sipencari-api/businesses/missings"
	"sipencari-api/businesses/users"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	likesmissingRepository _likeMissingMock.Repository
	likesmissingService    likesmissing.Usecase
	usersDomain            users.Domain
	missingsDomain         missings.Domain
	likesmissingDomain     likesmissing.Domain
)

func TestMain(m *testing.M) {
	likesmissingService = likesmissing.NewLikeMissingUsecase(&likesmissingRepository)
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
	likesmissingDomain = likesmissing.Domain{
		UserID: usersDomain.ID,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("TestGetAll | Valid", func(t *testing.T) {

		likesmissingRepository.On("GetAll", missingsDomain.ID).Return([]likesmissing.Domain{likesmissingDomain}).Once()
		result := likesmissingService.GetAll(missingsDomain.ID)
		assert.NotNil(t, result)
	})
	t.Run("TestGetAll | InValid", func(t *testing.T) {
		likesmissingRepository.On("GetAll", "").Return([]likesmissing.Domain{}).Once()
		result := likesmissingService.GetAll("")
		assert.NotNil(t, 0, len(result))
	})
}

func TestLike(t *testing.T) {
	t.Run("Like | Valid", func(t *testing.T) {
		likesmissingRepository.On("Like", usersDomain.ID, missingsDomain.ID, &likesmissingDomain).Return(likesmissing.Domain{}).Once()
		result := likesmissingService.Like(usersDomain.ID, missingsDomain.ID, &likesmissingDomain)
		assert.NotNil(t, result)
	})
	t.Run("Like | InValid", func(t *testing.T) {
		likesmissingRepository.On("Like", "", "", &likesmissing.Domain{}).Return(likesmissing.Domain{}).Once()
		result := likesmissingService.Like("", "", &likesmissing.Domain{})
		assert.NotNil(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {

		likesmissingRepository.On("GetByID", usersDomain.ID, missingsDomain.ID).Return(likesmissingDomain).Once()
		result := likesmissingService.GetByID(usersDomain.ID, missingsDomain.ID)
		assert.NotNil(t, result)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		likesmissingRepository.On("GetByID", "", "").Return(likesmissing.Domain{}).Once()
		result := likesmissingService.GetByID("", "")
		assert.NotNil(t, result)
	})
}

func TestUnlike(t *testing.T) {
	t.Run("Unlike | Valid", func(t *testing.T) {
		likesmissingRepository.On("Unlike", usersDomain.ID, missingsDomain.ID).Return(true).Once()
		result := likesmissingService.Unlike(usersDomain.ID, missingsDomain.ID)

		assert.NotNil(t, result)
	})

	t.Run("Unlike | InValid", func(t *testing.T) {
		likesmissingRepository.On("Unlike", "", "").Return(false).Once()

		result := likesmissingService.Unlike("", "")

		assert.NotNil(t, result)
	})
}
