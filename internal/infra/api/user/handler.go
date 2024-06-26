package user

import (
	"errors"
	"net/http"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService ports.UserService
}

func newHandler(service ports.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	user := &entity.User{}
	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}
	if err := u.userService.Create(user); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (u *userHandler) Login(c *gin.Context) {
	credentials := &entity.DefaultCredentials{}
	if err := c.Bind(credentials); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}
	token, err := u.userService.Login(credentials)

	if err != nil {
		println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Credentials not valid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (u *userHandler) Info(c *gin.Context) {
	// Get userID from token
	userID, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Credentials not valid",
		})
		return
	}
	user, err := u.userService.Info(userID)

	if err != nil {
		println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"name":     user.Name,
		"lastName": user.LastName,
		"email":    user.Email,
	})
}
