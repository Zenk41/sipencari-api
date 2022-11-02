package users

import (
	"net/http"
	"sipencari-api/app/middlewares"
	"sipencari-api/app/uploaders"
	"sipencari-api/businesses/users"
	"sipencari-api/controllers"
	"sipencari-api/controllers/users/request"
	"sipencari-api/controllers/users/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUsecase users.Usecase
}

func NewUserController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUsecase: authUC,
	}
}

func (ctrl *AuthController) GetAll(c echo.Context) error {
	usersData := ctrl.authUsecase.GetAll()

	users := []response.User{}

	for _, user := range usersData {
		if user.IsAdmin == false {
			users = append(users, response.FromDomain(user))
		}
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all users", users)
}

func (ctrl *AuthController) GetByID(c echo.Context) error {
	var id string = c.Param("id")
	user := ctrl.authUsecase.GetByID(id)

	if user.ID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "user found", response.FromDomain(user))
}

func (ctrl *AuthController) Register(c echo.Context) error {
	var result string
	picture, _ := c.FormFile("picture")
	if picture != nil {
		src, _ := picture.Open()
		defer src.Close()
		result, _ = uploaders.UploadToS3(c, picture.Filename, src)
	}
	input := request.User{}
	if result == "" {
		input.Picture = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	user := ctrl.authUsecase.Register(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "user created", response.FromDomain(user))
}

func (ctrl *AuthController) CreateAdmin(c echo.Context) error {
	var result string
	picture, _ := c.FormFile("picture")
	if picture != nil {
		src, _ := picture.Open()
		defer src.Close()
		result, _ = uploaders.UploadToS3(c, picture.Filename, src)
	}
	input := request.User{}
	if result == "" {
		input.Picture = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}
	user := ctrl.authUsecase.Register(input.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "admin created", response.FromDomain(user))
}

func (ctrl *AuthController) Update(c echo.Context) error {
	var id string = c.Param("id")

	var result string
	picture, _ := c.FormFile("picture")
	if picture != nil {
		src, _ := picture.Open()
		defer src.Close()
		result, _ = uploaders.UploadToS3(c, picture.Filename, src)
	}
	input := request.User{}
	if result == "" {
		input.Picture = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	user := ctrl.authUsecase.Update(id, input.ToDomain())

	if user.ID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "user updated", response.FromDomain(user))

}

func (ctrl *AuthController) Delete(c echo.Context) error {
	var id string = c.Param("id")

	if isSuccess := ctrl.authUsecase.Delete(id); !isSuccess {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user not found", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "user deleted", "")
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.UserLogin{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := ctrl.authUsecase.Login(userInput.ToDomain())
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (ctrl *AuthController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	if isListed := middlewares.CheckToken(user.Raw); !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}
	middlewares.Logout(user.Raw)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout succes",
	})
}
