package card

import (
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
	return s.repo.AddCard(card)
}

func (s *service) GetUserCard(cardID string, accountID string) ([]*entity.Card, error) {
	card, err := s.repo.GetUserCard(cardID, accountID, "email = ?", cardID)
	if err != nil {
		return nil, err
	}
	return card, nil
}
