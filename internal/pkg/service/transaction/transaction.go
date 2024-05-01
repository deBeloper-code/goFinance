package transaction

import (
	"errors"

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

// Transaction
func (s *service) TransactionUser(transaction *entity.Transaction) error {
	//Cards
	cardSource := &entity.Card{}
	cardDestination := &entity.Card{}
	// 1.- Find Source ID exists
	if err := s.repo.First(cardSource, "cardId=?", transaction.SourceCardID); err != nil {
		return err
	}

	// 2.- Find Destination ID exists
	if err := s.repo.First(cardDestination, "cardId=?", transaction.DestinationCardID); err != nil {
		return err
	}
	// 3.- Get balance from source user cardID
	//     Is the balance greater than or equal to the amount?
	balanceSource := cardSource.Balance
	if balanceSource < transaction.Amount {
		return errors.New("There is not enough balance.")
	}

	return nil
}
