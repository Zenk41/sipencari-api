package main

import (
	"encoding/json"
	"net/http"
	_middleware "sipencari-api/app/middlewares"
	_dbDriver "sipencari-api/drivers/mysql"
	"sipencari-api/drivers/mysql/users"
	"testing"

	_driverFactory "sipencari-api/drivers"

	_userUsecase "sipencari-api/businesses/users"
	_userController "sipencari-api/controllers/users"
	"sipencari-api/controllers/users/request"

	_categoryUsecase "sipencari-api/businesses/categories"
	_categoryController "sipencari-api/controllers/categories"

	_missingUsecase "sipencari-api/businesses/missings"
	_missingController "sipencari-api/controllers/missings"

	_commentUsecase "sipencari-api/businesses/comments"
	_commentController "sipencari-api/controllers/comments"

	_likeMissingUsecase "sipencari-api/businesses/likes_missing"
	_likeMissingController "sipencari-api/controllers/likes_missing"

	_likeCommentUsecase "sipencari-api/businesses/likes_comment"
	_likeCommentController "sipencari-api/controllers/likes_comment"

	_routes "sipencari-api/app/routes"
	util "sipencari-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func mainTest() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetEnv("DB_USERNAME"),
		DB_PASSWORD: util.GetEnv("DB_PASSWORD"),
		DB_HOST:     util.GetEnv("DB_HOST"),
		DB_PORT:     util.GetEnv("DB_PORT"),
		DB_NAME:     util.GetEnv("DB_TEST_NAME"),
	}
	db := configDB.InitDB()
	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:      util.GetEnv("JWT_SECRET_KEY"),
		ExpireDuration: 1,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUsecase)

	missingRepo := _driverFactory.NewMissingRepository(db)
	missingUsecase := _missingUsecase.NewMissingUsecase(missingRepo)
	missingCtrl := _missingController.NewMissingController(missingUsecase)

	commentRepo := _driverFactory.NewCommentRepository(db)
	commentUsecase := _commentUsecase.NewCommentUsecase(commentRepo)
	commentCtrl := _commentController.NewCommentController(commentUsecase)

	likeMissingRepo := _driverFactory.NewLikeMissingRepository(db)
	likeMissingUsecase := _likeMissingUsecase.NewLikeMissingUsecase(likeMissingRepo)
	likeMissingCtrl := _likeMissingController.NewLikeMissingController(likeMissingUsecase)

	likeCommentRepo := _driverFactory.NewLikeCommentRepository(db)
	likeCommentUsecase := _likeCommentUsecase.NewLikeCommentUsecase(likeCommentRepo)
	likeCommentCtrl := _likeCommentController.NewLikeCommentController(likeCommentUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:      configLogger.Init(),
		JWTMIddleware:         configJWT.Init(),
		AuthController:        *userCtrl,
		CategoryController:    *categoryCtrl,
		MissingController:     *missingCtrl,
		CommentController:     *commentCtrl,
		LikeMissingController: *likeMissingCtrl,
		LikeCommentController: *likeCommentCtrl,
	}

	routesInit.RouteRegister(e)
	return e
}

// Cleaning up seeds on database
func cleanUp(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode || http.StatusCreated == res.StatusCode {
		configDB := _dbDriver.ConfigDB{
			DB_USERNAME: util.GetEnv("DB_USERNAME"),
			DB_PASSWORD: util.GetEnv("DB_PASSWORD"),
			DB_HOST:     util.GetEnv("DB_HOST"),
			DB_PORT:     util.GetEnv("DB_PORT"),
			DB_NAME:     util.GetEnv("DB_TEST_NAME"),
		}
		db := configDB.InitDB()
		_dbDriver.CleanSeeds(db)
	}
}

func getJWTToken(t *testing.T) string {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetEnv("DB_USERNAME"),
		DB_PASSWORD: util.GetEnv("DB_PASSWORD"),
		DB_HOST:     util.GetEnv("DB_HOST"),
		DB_PORT:     util.GetEnv("DB_PORT"),
		DB_NAME:     util.GetEnv("DB_TEST_NAME"),
	}
	db := configDB.InitDB()
	user := _dbDriver.SeedUser(db)

	var userRequest *request.User = &request.User{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		Handler(mainTest()).
		Post("/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]
	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func getUser() users.User {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetEnv("DB_USERNAME"),
		DB_PASSWORD: util.GetEnv("DB_PASSWORD"),
		DB_HOST:     util.GetEnv("DB_HOST"),
		DB_PORT:     util.GetEnv("DB_PORT"),
		DB_NAME:     util.GetEnv("DB_TEST_NAME"),
	}

	db := configDB.InitDB()
	user := _dbDriver.SeedUser(db)

	return user
}

func TestRegister_Success(t *testing.T) {
	var userRequest *request.User = &request.User{
		Name:     "testing",
		Email:    "testing123@testing.com",
		Password: "testing123",
	}

	apitest.New().
		Observe(cleanUp).
		Handler(mainTest()).
		Post("/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *request.User = &request.User{
		Name:     "",
		Email:    "",
		Password: "",
	}
	apitest.New().
		Handler(mainTest()).
		Post("/signup").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

}

func TestLogin_Success(t *testing.T) {
	user := getUser()
	var userRequest *request.UserLogin = &request.UserLogin{
		Email:    user.Email,
		Password: user.Password,
	}
	apitest.New().
		Handler(mainTest()).
		Post("/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()

}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *request.UserLogin = &request.UserLogin{
		Email:    "",
		Password: "",
	}
	apitest.New().
		Handler(mainTest()).
		Post("/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

}

func TestLogin_Failed(t *testing.T) {
	var userRequest *request.UserLogin = &request.UserLogin{
		Email:    "doesntexist@gmail.com",
		Password: "itsfailed",
	}
	apitest.New().
		Handler(mainTest()).
		Post("/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusUnauthorized).
		End()
}


