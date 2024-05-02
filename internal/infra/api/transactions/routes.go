package transactions

import (
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/transaction"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := postgres.NewClient()
	service := transaction.NewService(repo)
	handler := newHandler(service)
	v1 := e.Group("/api/v1")
	// Create transaction
	v1.POST("/trantransaction", handler.CreateTransaction)
}
