package ports

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type CardRepository interface {
	AddCard(dest interface{}, value interface{}, conds ...interface{}) error
	Find(dest interface{}, conds ...interface{}) ([]map[string]interface{}, error)
}

type CardService interface {
	Add(card *entity.Card) error
	GetUserCard(userId string) ([]*entity.Card, error)
}
