package deposits

import (
	"errors"
	"net/http"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type depositHandler struct {
	depositService ports.DepositService
}

func newHandler(service ports.DepositService) *depositHandler {
	return &depositHandler{
		depositService: service,
	}
}

func (u *depositHandler) Add(c *gin.Context) {
	deposit := &entity.Deposit{}
	if err := c.Bind(deposit); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	// Get userID from token
	userID, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}
	deposit.UserId = userID

	if err := u.depositService.Add(deposit); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *depositHandler) GetDeposits(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	// Get userID from token
	userID, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}

	deposits, err := u.depositService.GetUserDeposit(userID, startDate, endDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, deposits)
}
