package routes

import (
	// "sipencari-api/app/middlewares"
	"sipencari-api/app/middlewares"
	"sipencari-api/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMIddleware    middleware.JWTConfig
	AuthController   users.AuthController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// logger
	e.Use(cl.LoggerMiddleware)

	e.POST("/signup", cl.AuthController.Register) // Registering user
	e.POST("/login", cl.AuthController.Login)     // Login

	admin := e.Group("", middleware.JWTWithConfig(cl.JWTMIddleware), middlewares.IsAdmin) // admin
	admin.POST("/users/admins", cl.AuthController.CreateAdmin)                            // Creating Admin

	users := e.Group("/users", middleware.JWTWithConfig(cl.JWTMIddleware))   // user
	users.GET("", cl.AuthController.GetAll)                                  // get all users
	users.GET("/:id", cl.AuthController.GetByID)                             // get user by ID
	users.PUT("/:id", cl.AuthController.Update, middlewares.IsAuthorized)    // update user by id
	users.DELETE("/:id", cl.AuthController.Delete, middlewares.IsAuthorized) // delete user

	withAuth := e.Group("", middleware.JWTWithConfig(cl.JWTMIddleware)) // with auth
	withAuth.POST("/logout", cl.AuthController.Logout)                  // logout

}
