package comments_test

import (
	"sipencari-api/businesses/comments"
	_commentMock "sipencari-api/businesses/comments/mocks"
	"sipencari-api/businesses/missings"
	"sipencari-api/businesses/users"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	commentsRepository _commentMock.Repository
	commentsService    comments.Usecase
	usersDomain        users.Domain
	missingsDomain     missings.Domain
	commentsDomain     comments.Domain
	commentDomain      []comments.Domain
)

func TestMain(m *testing.M) {
	commentsService = comments.NewCommentUsecase(&commentsRepository)
	usersDomain = users.Domain{
		ID:       "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613",
		Name:     "testing",
		Email:    "testing123@testing.com",
		Password: "testing123",
	}
	commentsDomain = comments.Domain{
		Message:   "saya pernah bertemu di jl cendrawsih",
		MissingID: missingsDomain.ID,
		UserID:    usersDomain.ID,
		UserName:  usersDomain.Name,
	}
	commentDomain = append(commentDomain, commentsDomain)
	missingsDomain = missings.Domain{
		ID:         "7a3a1ed7-b1c0-4079-a41a-f7253ddc2613",
		Title:      "Anak Hilang 5 Tahun Setelah Pulang Sekolah",
		Content:    "disini gua mau coba cari adik gua yang udah hilang beberapa tahun. terakhir ada dirumah itu sekitar tahun 2015/2016. kita udah coba cari kemana2, bahkan udah lapor polisi juga, tapi tetep g ada hasil nya.",
		CategoryID: 1,
		UserID:     usersDomain.ID,
		Comments:   commentDomain,
	}
	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("TestGetAll | Valid", func(t *testing.T) {

		commentsRepository.On("GetAll", missingsDomain.ID).Return([]comments.Domain{commentsDomain}).Once()
		result := commentsService.GetAll(missingsDomain.ID)
		assert.NotNil(t, result)
	})
	t.Run("TestGetAll | InValid", func(t *testing.T) {
		commentsRepository.On("GetAll", "").Return([]comments.Domain{}).Once()
		result := commentsService.GetAll("")
		assert.NotNil(t, 0, len(result))
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		commentsRepository.On("Create", usersDomain.ID, missingsDomain.ID, &commentsDomain).Return(comments.Domain{}).Once()
		result := commentsService.Create(usersDomain.ID, missingsDomain.ID, &commentsDomain)
		assert.NotNil(t, result)
	})
	t.Run("Register | InValid", func(t *testing.T) {
		commentsRepository.On("Create", "", "", &comments.Domain{}).Return(comments.Domain{}).Once()
		result := commentsService.Create("", "", &comments.Domain{})
		assert.NotNil(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {

		commentsRepository.On("GetByID", missingsDomain.ID, "1").Return(commentsDomain).Once()
		result := commentsService.GetByID(missingsDomain.ID, "1")
		assert.NotNil(t, result)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		commentsRepository.On("GetByID", "", "-1").Return(comments.Domain{}).Once()
		result := commentsService.GetByID("", "-1")
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		commentsRepository.On("Update", usersDomain.ID, missingsDomain.ID, "1", &commentsDomain).Return(comments.Domain{}).Once()
		result := commentsService.Update(usersDomain.ID, missingsDomain.ID, "1", &commentsDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		commentsRepository.On("Update", "", "", "-1", &comments.Domain{}).Return(comments.Domain{}).Once()

		result := commentsService.Update("", "", "-1", &comments.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		commentsRepository.On("Delete", missingsDomain.ID, "1").Return(true).Once()
		result := commentsService.Delete(missingsDomain.ID, "1")

		assert.NotNil(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		commentsRepository.On("Delete", "", "-1").Return(false).Once()

		result := commentsService.Delete("", "-1")

		assert.NotNil(t, result)
	})
}
