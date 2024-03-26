package ports

import "github.com/deBeloper-code/goFinance/internal/pkg/entity"

type AccountRepository interface {
	Add(value interface{}) error
	GetUserAccount(value interface{}, conds ...interface{}) error
}

type AccountService interface {
	Add(account *entity.Account) error
	GetUserAccount(userID, accountID string) error
}
