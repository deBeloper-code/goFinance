package user

import (
	"github.com/deBeloper-code/goFinance/internal/infra/api/middleware"
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := postgres.NewClient()
	service := user.NewService(repo)
	handler := newHandler(service)
	v1 := e.Group("/api/v1")
	// Create user
	v1.POST("/users", handler.CreateUser)
	// Login user
	v1.POST("/authentication", handler.Login)
	// Info user
	v1.GET("/user", middleware.Authenticate(), handler.Info)
}
