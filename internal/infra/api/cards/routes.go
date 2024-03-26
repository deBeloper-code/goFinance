package cards

import (
	"github.com/deBeloper-code/goFinance/internal/infra/api/middleware"
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/firestore"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/card"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := firestore.NewClient()
	service := card.NewService(repo)
	handler := newHandler(service)
	v1 := e.Group("/api/v1")
	// Add an card
	v1.POST("/cards", middleware.Authenticate(), handler.Add)
	v1.GET("/cards", middleware.Authenticate(), handler.GetCard)

}
