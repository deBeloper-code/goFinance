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

	// Add an expense
	e.POST("/api/v1/expenses", middleware.Authenticate(), handler.Add)
}
