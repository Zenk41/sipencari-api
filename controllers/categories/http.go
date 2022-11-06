package categories

import (
	"net/http"
	"sipencari-api/businesses/categories"
	"sipencari-api/controllers"
	"sipencari-api/controllers/categories/request"
	"sipencari-api/controllers/categories/response"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (ctrl *CategoryController) GetAll(c echo.Context) error {
	categoriesData := ctrl.categoryUseCase.GetAll()

	categories := []response.Category{}

	for _, category := range categoriesData {
		categories = append(categories, response.FromDomain(category))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (ctrl *CategoryController) GetByID(c echo.Context) error {
	var category_id string = c.Param("category_id")

	category := ctrl.categoryUseCase.GetByID(category_id)

	if category.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "category found", response.FromDomain(category))

}

func (ctrl *CategoryController) Create(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category := ctrl.categoryUseCase.Create(input.ToDomain())

	return controllers.NewResponse(c, http.StatusOK, "success", "category created", response.FromDomain(category))
}

func (ctrl *CategoryController) Update(c echo.Context) error {
	input := request.Category{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	var categoryID string = c.Param("category_id")

	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	category := ctrl.categoryUseCase.Update(categoryID, input.ToDomain())

	if category.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category created", response.FromDomain(category))
}

func (ctrl *CategoryController) Delete(c echo.Context) error {
	var categoryID string = c.Param("category_id")

	if isSuccess := ctrl.categoryUseCase.Delete(categoryID); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
