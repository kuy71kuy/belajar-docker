package route

import (
	"github.com/labstack/echo/middleware"
	"project/constants"
	"project/controller"
	m "project/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	///users route
	m.LogMiddleware(e)
	eNormal := e.Group("/users")
	eJwt := e.Group("/users")
	eJwt.Use(middleware.JWT([]byte(constants.LoadEnv("SECRET_KEY"))))
	eJwt.GET("", controller.GetUsersController)
	eJwt.GET("/:id", controller.GetUserController)
	eNormal.POST("", controller.CreateUserController)
	eNormal.POST("/login", controller.LoginUserController)
	eJwt.DELETE("/:id", controller.DeleteUserController)
	eJwt.PUT("/:id", controller.UpdateUserController)

	//books route
	eBook := e.Group("")
	eBook.Use(middleware.JWT([]byte(constants.LoadEnv("SECRET_KEY"))))
	eBook.GET("/books", controller.GetBooksController)
	eBook.GET("/books/:id", controller.GetBookController)
	eBook.POST("/books", controller.CreateBookController)
	eBook.DELETE("/books/:id", controller.DeleteBookController)
	eBook.PUT("/books/:id", controller.UpdateBookController)

	return e
}
