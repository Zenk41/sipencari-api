package routes

import (
	// "sipencari-api/app/middlewares"
	"sipencari-api/app/middlewares"
	"sipencari-api/controllers/categories"
	"sipencari-api/controllers/comments"
	"sipencari-api/controllers/likes_comment"
	"sipencari-api/controllers/likes_missing"
	"sipencari-api/controllers/missings"
	"sipencari-api/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware      echo.MiddlewareFunc
	JWTMIddleware         middleware.JWTConfig
	AuthController        users.AuthController
	CategoryController    categories.CategoryController
	MissingController     missings.MissingController
	CommentController     comments.CommentController
	LikeMissingController likesmissing.LikesMissingController
	LikeCommentController likescomment.LikesCommentController
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

	categories := e.Group("/categories", middleware.JWTWithConfig(cl.JWTMIddleware))      // category
	categories.GET("", cl.CategoryController.GetAll)                                      // get all category
	categories.GET("/:category_id", cl.CategoryController.GetByID)                        // get category by ID
	categories.POST("", cl.CategoryController.Create)                                     // create category
	categories.PUT("/:category_id", cl.CategoryController.Update, middlewares.IsAdmin)    // update category
	categories.DELETE("/:category_id", cl.CategoryController.Delete, middlewares.IsAdmin) // delete category

	missings := e.Group("/missings", middleware.JWTWithConfig(cl.JWTMIddleware)) // missing post
	missings.GET("", cl.MissingController.GetAll)                                // get all missings post
	missings.GET("/:missing_id", cl.MissingController.GetByID)                   // get category by ID
	missings.POST("", cl.MissingController.Create)                               // create missing post
	missings.PUT("/:missing_id", cl.MissingController.Update)                    // update missing
	missings.DELETE("/:missing_id", cl.MissingController.Delete)                 // delete missing

	comments := e.Group("/missings/:missing_id/comments", middleware.JWTWithConfig(cl.JWTMIddleware)) // comment
	comments.GET("", cl.CommentController.GetAll)                                                     // get all commments
	comments.GET("/:comment_id", cl.CommentController.GetByID)                                        // get commment by ID
	comments.POST("", cl.CommentController.Create)                                                    // create commment
	comments.PUT("/:comment_id", cl.CommentController.Update)                                         // update commment
	comments.DELETE("/:comment_id", cl.CommentController.Delete)                                      // delete commment

	likesmissing := e.Group("/missings/:missing_id/likes", middleware.JWTWithConfig(cl.JWTMIddleware)) // likes missing
	likesmissing.GET("", cl.LikeMissingController.GetAll)                                              // get all likesmissing
	likesmissing.GET("/:like_id", cl.LikeMissingController.GetByID)                                    // get likesmissing by ID
	likesmissing.POST("", cl.LikeMissingController.Like)                                               // like missing
	likesmissing.DELETE("", cl.LikeMissingController.Unlike)                                           // Unlike Missing

	likescomment := e.Group("/missings/:missing_id/comments/:comment_id/likes", middleware.JWTWithConfig(cl.JWTMIddleware)) // likes comment
	likescomment.GET("", cl.LikeCommentController.GetAll)                                                                   // get all likescomment
	likescomment.GET("/:like_id", cl.LikeCommentController.GetByID)                                                         // get likescomment by ID
	likescomment.POST("", cl.LikeCommentController.Like)                                                                    // create likescomment
	likescomment.DELETE("", cl.LikeCommentController.Unlike)                                                                // unlike comment

	withAuth := e.Group("", middleware.JWTWithConfig(cl.JWTMIddleware)) // with auth
	withAuth.POST("/logout", cl.AuthController.Logout)                  // logout

}
