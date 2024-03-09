package expenses

import (
	"errors"
	"net/http"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type expenseHandler struct {
	expenseService ports.ExpenseService
}

func newHandler(service ports.ExpenseService) *expenseHandler {
	return &expenseHandler{
		expenseService: service,
	}
}

func (u *expenseHandler) Add(c *gin.Context) {
	expense := &entity.Expense{}
	if err := c.Bind(expense); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	// Get userID from token
	userID, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}
	expense.UserId = userID

	if err := u.expenseService.Add(expense); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	c.JSON(http.StatusOK, nil)
}

// func (u *expenseHandler) GetUserexpense(c *gin.Context) {

// 	if err := c.Bind(credentials); err != nil {
// 		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
// 		return
// 	}
// 	token, err := u.ExpenseService.GetUserexpense(credentials)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, nil)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": token,
// 	})
// }
