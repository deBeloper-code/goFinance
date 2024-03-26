package ports

import "github.com/deBeloper-code/goFinance/internal/pkg/entity"

type AccountRepository interface {
	Add(account *entity.Account) error
	GetUserAccount(userID, accountID string) (*entity.Account, error)
}

type AccountService interface {
	Add(account *entity.Account) error
	GetUserAccount(userID, accountID string) (*entity.Account, error)
}
