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

func newHandler(service ports.CardService) *cardHandler {
	return &cardHandler{
		cardService: service,
	}
}

func (u *cardHandler) Add(c *gin.Context) {
	var card entity.Card

	// Checking json body
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Sending data to (Business Rules)
	if err := u.cardService.Add(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Success response
	c.JSON(http.StatusOK, gin.H{"Success": "Card cread"})
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
