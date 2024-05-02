package transaction

import (
	"errors"

	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
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

	// 4.- Start transaction process.
	// 4.1- Subtract BALANCE - AMOUNT
	updatedBalanceSource := balanceSource - transaction.Amount
	// 4.2- Sum BALANCE + AMOUNT
	updatedBalanceDestination := cardDestination.Balance + transaction.Amount
	// 4.3- Fill values to send
	paramsSource := postgres.SourceParams{}
	paramsDestination := postgres.DestinationParams{}

	transactionParams := postgres.TransactionParams{
		transactionInit: "",
		source:          paramsSource,
		destination:     paramsDestination,
	}
	// 4.4- Sending Transaction
	if err := s.repo.SendTransaction(transactionParams); err != nil {
		return err
	}

	// 5.- MISSION COMPLETED
	return nil
}
