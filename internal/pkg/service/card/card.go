package card

import (
	"time"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
)

type service struct {
	repo ports.CardRepository
}

func NewService(repo ports.CardRepository) *service {
	return &service{
		repo: repo,
	}
}

// This layer is for adding Business Rules
func (s *service) Add(card *entity.Card) error {
	// User table
	user := &entity.User{}

	return s.repo.AddCard(card, user, "id = ?", card.UserID)
}

func (s *service) GetUserCard(userId string) ([]*entity.Card, error) {
	// Cards slice
	var cards []*entity.Card
	// Retrieving cards from repository
	cardInterfaces, err := s.repo.Find(&entity.Card{}, "user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	for _, m := range cardInterfaces {
		card := entity.Card{
			CardID:        m["card_id"].(string),
			UserID:        m["user_id"].(string),
			AccountNumber: m["account_number"].(string),
			CardNumber:    m["card_number"].(string),
			Balance:       m["balance"].(float64),
			// TODO: We need to review the dates UTC
			ExpiryDate: m["expiry_date"].(time.Time),
			CVV:        m["cvv"].(string),
			Issuer:     m["issuer"].(string),
			Type:       m["type"].(string),
		}
		cards = append(cards, &card)
	}

	return cards, nil
}
