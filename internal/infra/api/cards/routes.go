package cards

import (
	"github.com/deBeloper-code/goFinance/internal/infra/api/middleware"
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/card"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// DB
	repo := postgres.NewClient()
	// Service (Business Rules)
	service := card.NewService(repo)
	// Handler (Response/Request)
	handler := newHandler(service)
	// Group routes
	v1 := e.Group("/api/v1")
	// Add an card
	v1.POST("/cards", handler.Add)
	// Get a user card
	v1.GET("/cards", middleware.Authenticate(), handler.GetCard)

}
