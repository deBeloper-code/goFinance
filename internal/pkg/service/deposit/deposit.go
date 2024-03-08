package deposit

import (
	"time"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/deBeloper-code/goFinance/internal/pkg/utils"
)

type service struct {
	repo ports.DepositRepository
}

func NewService(repo ports.DepositRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Add(deposit *entity.Deposit) error {
	return s.repo.AddDeposit(deposit)
}
func (s *service) GetUserDeposit(userID string, startDate, endDate string) ([]*entity.Deposit, error) {
	var st, ed time.Time
	var err error
	// 1. Date interval will be correct
	if !utils.IsDateRangeValid(startDate, endDate) {
		// 1.1 Generate date range by default
		st, ed = utils.GetLastSevenDays()
	} else {
		// 1.2 If the dates are valid We only parse to date time format
		st, ed, err = utils.ParseDateRange(startDate, endDate)
	}

	deposit, err := s.repo.GetUserDeposit(userID, st, ed)
	if err != nil {
		return nil, err
	}
	return deposit, nil
}
