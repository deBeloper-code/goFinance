package account

import (
	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/account"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// DataBase
	repo := postgres.NewClient()
	account.NewService(repo)
}
