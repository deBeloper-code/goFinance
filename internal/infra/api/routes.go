package api

import (
	"github.com/deBeloper-code/goFinance/internal/infra/api/account"
	"github.com/deBeloper-code/goFinance/internal/infra/api/deposits"
	"github.com/deBeloper-code/goFinance/internal/infra/api/expenses"
	"github.com/deBeloper-code/goFinance/internal/infra/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// User
	user.RegisterRoutes(e)
	// Account
	account.RegisterRoutes(e)
	// Deposits
	deposits.RegisterRoutes(e)
	// Expense
	expenses.RegisterRoutes(e)
}
