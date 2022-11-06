package likescomment_test

import (
	"sipencari-api/businesses/comments"
	likescomment "sipencari-api/businesses/likes_comment"
	_likesCommentMock "sipencari-api/businesses/likes_comment/mocks"
	"sipencari-api/businesses/missings"
	"sipencari-api/businesses/users"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	likescommentRepository _likesCommentMock.Repository
	likescommentService    likescomment.Usecase
	usersDomain            users.Domain
	missingsDomain         missings.Domain
	likescommentDomain     likescomment.Domain
	commentsDomain         comments.Domain
	commentDomain          []comments.Domain
)

func TestMain(m *testing.M) {
	likescommentService = likescomment.NewLikeCommentUsecase(&likescommentRepository)
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
	likescommentDomain = likescomment.Domain{
		UserID: usersDomain.ID,
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("TestGetAll | Valid", func(t *testing.T) {

		likescommentRepository.On("GetAll", missingsDomain.ID, "1").Return([]likescomment.Domain{likescommentDomain}).Once()
		result := likescommentService.GetAll(missingsDomain.ID, "1")
		assert.NotNil(t, result)
	})
	t.Run("TestGetAll | InValid", func(t *testing.T) {
		likescommentRepository.On("GetAll", "", "").Return([]likescomment.Domain{}).Once()
		result := likescommentService.GetAll("", "")
		assert.NotNil(t, 0, len(result))
	})
}

func TestLike(t *testing.T) {
	t.Run("Like | Valid", func(t *testing.T) {
		likescommentRepository.On("Like", usersDomain.ID, missingsDomain.ID, 1, &likescommentDomain).Return(likescomment.Domain{}).Once()
		result := likescommentService.Like(usersDomain.ID, missingsDomain.ID, 1, &likescommentDomain)
		assert.NotNil(t, result)
	})
	t.Run("Like | InValid", func(t *testing.T) {
		likescommentRepository.On("Like", "", "", -1, &likescomment.Domain{}).Return(likescomment.Domain{}).Once()
		result := likescommentService.Like("", "", -1, &likescomment.Domain{})
		assert.NotNil(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {

		likescommentRepository.On("GetByID", usersDomain.ID, missingsDomain.ID, 1).Return(likescommentDomain).Once()
		result := likescommentService.GetByID(usersDomain.ID, missingsDomain.ID, 1)
		assert.NotNil(t, result)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		likescommentRepository.On("GetByID", "", "", -1).Return(likescomment.Domain{}).Once()
		result := likescommentService.GetByID("", "", -1)
		assert.NotNil(t, result)
	})
}

func TestUnlike(t *testing.T) {
	t.Run("Unlike | Valid", func(t *testing.T) {
		likescommentRepository.On("Unlike", usersDomain.ID, missingsDomain.ID, 1).Return(true).Once()
		result := likescommentService.Unlike(usersDomain.ID, missingsDomain.ID, 1)

		assert.NotNil(t, result)
	})

	t.Run("Unlike | InValid", func(t *testing.T) {
		likescommentRepository.On("Unlike", "", "", -1).Return(false).Once()

		result := likescommentService.Unlike("", "", -1)

		assert.NotNil(t, result)
	})
}
