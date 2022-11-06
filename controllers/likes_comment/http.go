package likescomment

import (
	"net/http"
	"sipencari-api/app/middlewares"
	likescomment "sipencari-api/businesses/likes_comment"
	"sipencari-api/controllers"
	"sipencari-api/controllers/likes_comment/request"
	"sipencari-api/controllers/likes_comment/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikesCommentController struct {
	likeCommentUsecase likescomment.Usecase
}

func NewLikeCommentController(likeCommentUC likescomment.Usecase) *LikesCommentController {
	return &LikesCommentController{
		likeCommentUsecase: likeCommentUC,
	}
}

func (ctrl *LikesCommentController) GetAll(c echo.Context) error {
	var missingId string = c.Param("missing_id")
	commentId := c.Param("comment_id")
	likeData := ctrl.likeCommentUsecase.GetAll(missingId, commentId)
	likeComments := []response.LikeComment{}
	for _, like := range likeData {
		likeComments = append(likeComments, response.FromDomain(like))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all likes missing", likeComments)
}

func (ctrl *LikesCommentController) GetByID(c echo.Context) error {
	var id string = c.Param("like_id")
	var missingId string = c.Param("missing_id")
	commentId, _ := strconv.Atoi(c.Param("comment_id"))
	likeComment := ctrl.likeCommentUsecase.GetByID(id, missingId, commentId)
	if likeComment.UserID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like in comment not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "like in comment found", response.FromDomain(likeComment))
}

func (ctrl *LikesCommentController) Like(c echo.Context) error {
	userID := middlewares.GetUserID(c)
	input := request.LikeComment{}
	var missingID string = c.Param("missing_id")
	commentId, _ := strconv.Atoi(c.Param("comment_id"))
	input.UserID = userID
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	likeComment := ctrl.likeCommentUsecase.Like(userID, missingID, commentId, input.ToDomain())

	if likeComment.UserID != userID {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like comment not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "liked comment", response.FromDomain(likeComment))
}

func (ctrl *LikesCommentController) Unlike(c echo.Context) error {
	userID := middlewares.GetUserID(c)
	var missingID string = c.Param("missing_id")
	commentId, _ := strconv.Atoi(c.Param("comment_id"))

	if isSuccess := ctrl.likeCommentUsecase.Unlike(userID, missingID, commentId); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "cannot unlike like not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "unliked comment", "")
}
