package account

import (
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/account"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// DataBase
	repo := postgres.NewClient()
	// Service
	service := account.NewService(repo)
	// Handler
	handler := NewHandler(service)
	// Routes
	v1 := e.Group("/api/v1")
	// Create user
	v1.POST("/accounts", handler.Add)
	// Login user
	v1.GET("/accounts", handler.GetUserAccount)
}
