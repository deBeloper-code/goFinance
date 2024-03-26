package account

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
)

type service struct {
	repo ports.AccountRepository
}

// This generate a new instance
func NewService(repo ports.AccountRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Add(account *entity.Account) error {
	// TODO: Generate Logic for example, generate Account number
	return s.repo.Add(account)
}

func (s *service) GetUserAccount(userID, accountID string) (*entity.Account, error) {
	return s.repo.GetUserAccount(userID, accountID)
}
