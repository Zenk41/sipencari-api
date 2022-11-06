package missings

import (
	"net/http"
	"sipencari-api/app/middlewares"
	"sipencari-api/app/uploaders"
	"sipencari-api/businesses/missings"
	"sipencari-api/controllers"
	"sipencari-api/controllers/missings/request"
	"sipencari-api/controllers/missings/response"

	"github.com/labstack/echo/v4"
)

type MissingController struct {
	missingUsecase missings.Usecase
}

func NewMissingController(missingUC missings.Usecase) *MissingController {
	return &MissingController{
		missingUsecase: missingUC,
	}
}

func (ctrl *MissingController) GetAll(c echo.Context) error {
	missingData := ctrl.missingUsecase.GetAll()

	missings := []response.Missing{}

	for _, missing := range missingData {
		missings = append(missings, response.FromDomain(missing))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all missing", missings)
}

func (ctrl *MissingController) GetByID(c echo.Context) error {
	var id string = c.Param("missing_id")
	missing := ctrl.missingUsecase.GetByID(id)
	if missing.ID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "missing not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "missing post found", response.FromDomain(missing))
}

func (ctrl *MissingController) Create(c echo.Context) error {
	var result string
	userID := middlewares.GetUserID(c)

	input := request.Missing{}

	image, _ := c.FormFile("image")
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = uploaders.UploadToS3(c, image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	missing := ctrl.missingUsecase.Create(userID, input.ToDomain())

	return controllers.NewResponse(c, http.StatusOK, "success", "missing post created", response.FromDomain(missing))

}

func (ctrl *MissingController) Update(c echo.Context) error {
	var result string
	input := request.Missing{}
	userID := middlewares.GetUserID(c)
	var missingID string = c.Param("missing_id")
	image, _ := c.FormFile("image")
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = uploaders.UploadToS3(c, image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	missing := ctrl.missingUsecase.Update(userID, missingID, input.ToDomain())

	if missing.ID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "missing post not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "missing post updated", response.FromDomain(missing))
}

func (ctrl *MissingController) Delete(c echo.Context) error {
	var missingID string = c.Param("missing_id")
	missing := ctrl.missingUsecase.GetByID(missingID)
	UserID := middlewares.GetUserID(c)
	if UserID != missing.UserID {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "Unauthorized", "")
	}

	if isSuccess := ctrl.missingUsecase.Delete(missingID); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "missing post not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "missing post deleted", "")
}
