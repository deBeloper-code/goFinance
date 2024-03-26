package transaction

import (
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
)

type service struct {
	repo ports.TransactionRepository
}

func NewService(repo ports.TransactionRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Add(transaction *entity.Transaction) error {
	return s.repo.Add(transaction)
}

func (s *service) GetUserCard(cardID, accountID string, optionalParams ...string) ([]*entity.Transaction, error) {
	var typeOfTransaction string
	if len(optionalParams) > 0 {
		typeOfTransaction = optionalParams[0]
	}
	transaction, err := s.repo.GetTransaction(cardID, accountID, typeOfTransaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
