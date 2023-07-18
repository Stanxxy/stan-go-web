package controller

import (
	"net/http"

	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/Stanxxy/stan-go-web/internal/models"
	"github.com/Stanxxy/stan-go-web/internal/models/user"
	echo "github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

// Only for testing
func (ctrl User) AddUser(c echo.Context) error {

	cc := c.(*context.AppContext)

	var user models.User

	err := cc.Bind(&user)

	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	err = cc.UserStore.Create(&user)

	if err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	return c.String(http.StatusOK, "User created")
}

// Only for testing
func (ctrl User) RemoveUser(c echo.Context) error {

	cc := c.(*context.AppContext)

	var user models.User

	err := cc.Bind(&user)

	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	err = cc.UserStore.Delete(&user)

	if err != nil {
		b := errors.NewBoom(errors.EntityDeleteError, errors.ErrorText(errors.EntityDeleteError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	return c.String(http.StatusOK, "User removed")
}

// Only for testing
func (ctrl User) UpdateUser(c echo.Context) error {

	cc := c.(*context.AppContext)

	var user models.User

	err := cc.Bind(&user)

	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	err = cc.UserStore.Save(&user)

	if err != nil {
		b := errors.NewBoom(errors.EntityUpdateError, errors.ErrorText(errors.EntityUpdateError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// Do something with the user object
	return c.String(http.StatusOK, "User updated")
}

// Only for testing
func (ctrl UserList) GetUsers(c echo.Context) error {
	cc := c.(*context.AppContext)

	users := []models.User{}

	err := cc.UserStore.RetrieveMany(&users)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}
	return c.JSON(http.StatusOK, users)
}

// Only for testing
func (ctrl User) GetUser(c echo.Context) error {
	cc := c.(*context.AppContext)
	userID, err := xid.FromString(c.Param("id"))

	if err != nil {
		b := errors.NewBoom(errors.TypeConvertionError, errors.ErrorText(errors.TypeConvertionError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	user := models.User{ID: userID}

	err = cc.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	return c.JSON(http.StatusOK, user)
}
