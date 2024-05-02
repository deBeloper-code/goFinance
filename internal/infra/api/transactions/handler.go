package transactions

import (
	"errors"
	"net/http"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type transactionsHandler struct {
	transactionsService ports.TransactionService
}

func newHandler(service ports.TransactionService) *transactionsHandler {
	return &transactionsHandler{
		transactionsService: service,
	}
}

func (u *transactionsHandler) CreateTransaction(c *gin.Context) {
	transaction := &entity.Transaction{}
	if err := c.Bind(transaction); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}
	if err := u.transactionsService.TransactionUser(transaction); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)

}
