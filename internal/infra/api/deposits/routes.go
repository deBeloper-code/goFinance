package deposits

import (
	"github.com/deBeloper-code/goFinance/internal/infra/api/middleware"
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/firestore"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/deposit"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := firestore.NewClient()
	service := deposit.NewService(repo)
	handler := newHandler(service)
	v1 := e.Group("/api/v1")
	// Add an deposit
	v1.POST("/deposits", middleware.Authenticate(), handler.Add)
	v1.GET("/deposits", middleware.Authenticate(), handler.GetDeposits)

}
