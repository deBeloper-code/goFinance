package ports

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type TransactionRepository interface {
	First(dest interface{}, conds ...interface{}) error
	SendTransaction(transactionInit interface{}, source interface{}, scolumn string, svalue interface{}, destination interface{}, dcolumn string, dvalue interface{}) error
}

type TransactionService interface {
	TransactionUser(transaction *entity.Transaction) error
}
