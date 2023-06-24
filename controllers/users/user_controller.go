package users

import (
	"github.com/gin-gonic/gin"
	"mypackage/domain/users"
	"mypackage/services"
	"mypackage/utils/errors"
	"net/http"
	"strconv"
	"time"
)

func CreateUser(c *gin.Context) {
	var newUser users.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		// Unable to bind request body to User JSON. Case of bad request
		restErr := errors.NewBadRequestError("Unable to bind request body to User JSON")
		c.JSON(restErr.Status, restErr)
		return
	}
	currentTime := time.Now()
	timestampString := currentTime.Format("2006-01-02 15:04:05")
	newUser.DateCreated = timestampString

	// Passing the bound data object to the appropriate service
	result, saveErr := services.CreateNewUser(newUser)
	if saveErr != nil {
		// Handle create new user error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	// Prepare data to be used from the request
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id is not integer")
		c.JSON(restErr.Status, restErr)
		return
	}

	var currentUser users.User
	currentUser.ID = userId
	result, getErr := services.GetUserByID(currentUser)
	if getErr != nil {
		// Handle unable to retrieve
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var userToUpdate users.User
	err := c.ShouldBindJSON(&userToUpdate)
	if err != nil {
		// Unable to bind request body to User JSON. Case of bad request
		restErr := errors.NewBadRequestError("Unable to bind request body to User JSON")
		c.JSON(restErr.Status, restErr)
		return
	}

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id is not integer")
		c.JSON(restErr.Status, restErr)
		return
	}

	userToUpdate.ID = userId

	result, getErr := services.UpdateUserByID(userToUpdate)
	if getErr != nil {
		// Handle unable to retrieve
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("user_id is not integer")
		c.JSON(restErr.Status, restErr)
		return
	}

	var userToDelete users.User
	userToDelete.ID = userId

	result, getErr := services.DeleteUserByID(userToDelete) // NOTE : result should be the deleted user
	if getErr != nil {
		// Handle unable to retrieve
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
