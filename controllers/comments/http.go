package comments

import (
	"net/http"
	"sipencari-api/app/middlewares"
	"sipencari-api/app/uploaders"
	"sipencari-api/businesses/comments"
	"sipencari-api/controllers"
	"sipencari-api/controllers/comments/request"
	"sipencari-api/controllers/comments/response"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	commentUsecase comments.Usecase
}

func NewCommentController(commentUC comments.Usecase) *CommentController {
	return &CommentController{
		commentUsecase: commentUC,
	}
}

func (ctrl *CommentController) GetAll(c echo.Context) error {
	missingID := c.Param("missing_id")
	commentData := ctrl.commentUsecase.GetAll(missingID)

	comments := []response.Comment{}

	for _, comment := range commentData {
		comments = append(comments, response.FromDomain(comment))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all comment", comments)
}

func (ctrl *CommentController) GetByID(c echo.Context) error {
	var id string = c.Param("comment_id")
	missingID := c.Param("missing_id")
	comment := ctrl.commentUsecase.GetByID(missingID, id)

	if comment.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "comment not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "comment post found", response.FromDomain(comment))
}

func (ctrl *CommentController) Create(c echo.Context) error {
	var result string
	userID := middlewares.GetUserID(c)
	var missingID string = c.Param("missing_id")
	input := request.Comment{}

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

	comment := ctrl.commentUsecase.Create(userID, missingID, input.ToDomain())

	return controllers.NewResponse(c, http.StatusOK, "success", "comment created", response.FromDomain(comment))

}

func (ctrl *CommentController) Update(c echo.Context) error {
	var result string
	input := request.Comment{}
	userID := middlewares.GetUserID(c)
	var missingID string = c.Param("missing_id")
	var commentID string = c.Param("comment_id")
	commentFind := ctrl.commentUsecase.GetByID(missingID, commentID)
	if userID != commentFind.UserID {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "Unauthorized", "")
	}
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

	comment := ctrl.commentUsecase.Update(userID, missingID, commentID, input.ToDomain())

	if comment.ID == 0 {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "comment not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "comment updated", response.FromDomain(comment))
}

func (ctrl *CommentController) Delete(c echo.Context) error {
	var commentID string = c.Param("comment_id")
	var missingID string = c.Param("missing_id")
	comment := ctrl.commentUsecase.GetByID(missingID, commentID)
	userID := middlewares.GetUserID(c)
	if userID != comment.UserID {
		return controllers.NewResponse(c, http.StatusUnauthorized, "failed", "Unauthorized", "")
	}

	if isSuccess := ctrl.commentUsecase.Delete(missingID,commentID); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "comment not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "comment deleted", "")
}
