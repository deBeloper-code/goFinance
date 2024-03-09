package expenses

import (
	"github.com/deBeloper-code/goFinance/internal/infra/api/middleware"
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/firestore"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/expense"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := firestore.NewClient()
	service := expense.NewService(repo)
	handler := newHandler(service)
	v1 := e.Group("/api/v1")

	// Add an expense
	v1.POST("/expenses", middleware.Authenticate(), handler.Add)
	v1.GET("/expenses", middleware.Authenticate(), handler.GetExpenses)
}
