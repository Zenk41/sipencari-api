package likesmissing

import (
	"net/http"
	"sipencari-api/app/middlewares"
	likesmissing "sipencari-api/businesses/likes_missing"
	"sipencari-api/controllers"
	"sipencari-api/controllers/likes_missing/request"
	"sipencari-api/controllers/likes_missing/response"

	"github.com/labstack/echo/v4"
)

type LikesMissingController struct {
	likeMissingUsecase likesmissing.Usecase
}

func NewLikeMissingController(likeMissingUC likesmissing.Usecase) *LikesMissingController {
	return &LikesMissingController{
		likeMissingUsecase: likeMissingUC,
	}
}

func (ctrl *LikesMissingController) GetAll(c echo.Context) error {
	missingID := c.Param("missing_id")
	likeData := ctrl.likeMissingUsecase.GetAll(missingID)
	likeMissings := []response.LikeMissing{}
	for _, like := range likeData {
		likeMissings = append(likeMissings, response.FromDomain(like))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all likes missing", likeMissings)
}
func (ctrl *LikesMissingController) GetByID(c echo.Context) error {
	var id string = c.Param("like_id")
	var missingId string = c.Param("missing_id")
	likeMissing := ctrl.likeMissingUsecase.GetByID(id, missingId)
	if likeMissing.UserID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like in missing not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "like in missing found", response.FromDomain(likeMissing))
}

func (ctrl *LikesMissingController) Like(c echo.Context) error {
	userID := middlewares.GetUserID(c)
	input := request.LikeMissing{}
	var missingID string = c.Param("missing_id")
	input.UserID = userID

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	likeMissing := ctrl.likeMissingUsecase.Like(userID, missingID, input.ToDomain())

	if likeMissing.UserID != userID {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like missing not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "liked missing", response.FromDomain(likeMissing))
}

func (ctrl *LikesMissingController) Unlike(c echo.Context) error {
	userID := middlewares.GetUserID(c)
	var missingID string = c.Param("missing_id")

	if isSuccess := ctrl.likeMissingUsecase.Unlike(userID, missingID); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "cannot unlike like not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "unliked missing", "")
}
