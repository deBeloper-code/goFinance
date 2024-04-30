package cards

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
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
	c.JSON(http.StatusOK, gin.H{"Success": "Card created"})
}

type CardInfoRes struct {
	CardID        string    `json:"cardId"`
	UserID        string    `json:"userId"`
	AccountNumber string    `json:"accountNumber"`
	CardNumber    string    `json:"cardNumber"`
	Balance       float64   `json:"balance"`
	ExpiryDate    time.Time `json:"expiryDate"`
	CVV           string    `json:"cvv"`
	Type          string    `json:"type"`
}

func (u *cardHandler) GetCard(c *gin.Context) {
	// Get userID from token
	userId, ok := c.MustGet("userID").(string)
	if !ok {
		c.JSON(http.StatusBadRequest, errors.New("Invalid token"))
		return
	}

	cards, err := u.cardService.GetUserCard(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var cardInfo []CardInfoRes

	for _, card := range cards {
		cardInfo = append(cardInfo, CardInfoRes{
			CardID:        card.CardID,
			UserID:        card.UserID,
			AccountNumber: card.AccountNumber,
			CardNumber:    card.CardNumber,
			Balance:       card.Balance,
			ExpiryDate:    card.ExpiryDate,
			CVV:           card.CVV,
			Type:          card.Type,
		})
	}

	c.JSON(http.StatusOK, cardInfo)
}
