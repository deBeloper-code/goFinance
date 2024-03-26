package ports

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type CardRepository interface {
	Add(card *entity.Card) error
	GetUserCard(cardID string, accountID string) ([]*entity.Card, error)
}

type CardService interface {
	Add(card *entity.Card) error
	GetUserCard(cardID string, accountID string) ([]*entity.Card, error)
}
