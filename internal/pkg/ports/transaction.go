package ports

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type TransactionRepository interface {
	Add(card *entity.Transaction) error
	GetTransaction(transactionID, accountID string, optionalParams ...string) ([]*entity.Transaction, error)
}

type TransactionService interface {
	Add(transaction *entity.Transaction) error
	GetTransaction(transactionID, accountID string, optionalParams ...string) ([]*entity.Transaction, error)
}
