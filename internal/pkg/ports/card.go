package ports

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type CardRepository interface {
	AddCard(dest interface{}) error
	GetUserCard(cardID, accountID string, conds ...interface{}) ([]*entity.Card, error)
}

type CardService interface {
	Add(card *entity.Card) error
	GetUserCard(cardID string, accountID string) ([]*entity.Card, error)
}
