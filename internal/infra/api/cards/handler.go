package cards

import (
	"errors"
	"net/http"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type cardHandler struct {
	cardService ports.CardService
}

func newHandler(service ports.CardRepository) *cardHandler {
	return &cardHandler{
		cardService: service,
	}
}

func (u *cardHandler) Add(c *gin.Context) {
	card := &entity.Card{}
	if err := c.Bind(card); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	// Get userID from token
	userID, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}
	card.AccountID = userID

	if err := u.cardService.Add(card); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid input"))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *cardHandler) GetCard(c *gin.Context) {
	cardID := c.Query("cardId")
	accountID := c.Query("accountId")
	// Get userID from token
	_, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}

	cards, err := u.cardService.GetUserCard(cardID, accountID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, cards)
}
