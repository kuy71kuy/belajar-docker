package controller

import (
	"net/http"
	"project/config"
	"project/middleware"
	"project/model"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})

}

func GetUserController(c echo.Context) error {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := model.User{}
	foundUser := model.User{}
	c.Bind(&user)
	user.Id = id
	if err := config.DB.First(&foundUser, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{user.Id, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}

//=========================================================================
//=========================================================================
//=========================================================================
//=========================================================================

func GetBooksController(c echo.Context) error {
	var books []model.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})

}

func GetBookController(c echo.Context) error {
	var book model.Book
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)
	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := model.Book{}
	foundBook := model.Book{}
	c.Bind(&book)
	book.Id = id
	if err := config.DB.First(&foundBook, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	var book model.Book
	id, _ := strconv.Atoi(c.Param("id"))
	if err := config.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}
