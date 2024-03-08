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

	// Add an deposit
	e.POST("/api/v1/deposits", middleware.Authenticate(), handler.Add)

}
