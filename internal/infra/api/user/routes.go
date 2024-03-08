package user

import (
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := postgres.NewClient()
	service := user.NewService(repo)
	handler := newHandler(service)

	// Create user
	e.POST("/api/v1/users", handler.CreateUser)
	// Login user
	e.POST("/api/v1/authentication", handler.Login)
}
