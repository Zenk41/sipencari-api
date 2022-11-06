package categories_test

import (
	// "sipencari-api/businesses/hashtags"
	// locations_category "sipencari-api/businesses/locations_category"
	"sipencari-api/businesses/categories"
	_categoryMock "sipencari-api/businesses/categories/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoriesRepository _categoryMock.Repository
	categoriesService    categories.Usecase

	categoriesDomain     categories.Domain
)

func TestMain(m *testing.M) {
	categoriesService = categories.NewCategoryUsecase(&categoriesRepository)
	categoriesDomain = categories.Domain{
		Name: "Binatang",
	}
	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("TestGetAll | Valid", func(t *testing.T) {

		categoriesRepository.On("GetAll").Return([]categories.Domain{categoriesDomain}).Once()
		result := categoriesService.GetAll()
		assert.NotNil(t, result)
	})
	t.Run("TestGetAll | InValid", func(t *testing.T) {
		categoriesRepository.On("GetAll").Return([]categories.Domain{}).Once()
		result := categoriesService.GetAll()
		assert.NotNil(t, 0, len(result))
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoriesRepository.On("Create", &categoriesDomain).Return(categories.Domain{}).Once()
		result := categoriesService.Create(&categoriesDomain)
		assert.NotNil(t, result)
	})
	t.Run("Register | InValid", func(t *testing.T) {
		categoriesRepository.On("Create", &categories.Domain{}).Return(categories.Domain{}).Once()
		result := categoriesService.Create(&categories.Domain{})
		assert.NotNil(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {

		categoriesRepository.On("GetByID", "1").Return(categoriesDomain).Once()
		result := categoriesService.GetByID("1")
		assert.NotNil(t, result)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		categoriesRepository.On("GetByID", "-1").Return(categories.Domain{}).Once()
		result := categoriesService.GetByID("-1")
		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoriesRepository.On("Update", "1", &categoriesDomain).Return(categories.Domain{}).Once()
		result := categoriesService.Update("1", &categoriesDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		categoriesRepository.On("Update", "-1", &categories.Domain{}).Return(categories.Domain{}).Once()

		result := categoriesService.Update("-1", &categories.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoriesRepository.On("Delete", "1").Return(true).Once()
		result := categoriesService.Delete("1")

		assert.NotNil(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		categoriesRepository.On("Delete", "-1").Return(false).Once()

		result := categoriesService.Delete("-1")

		assert.NotNil(t, result)
	})
}
