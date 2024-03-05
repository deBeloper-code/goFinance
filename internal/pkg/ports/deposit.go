package ports

import (
	"time"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type DepositRepository interface {
	AddDeposit(deposit *entity.Deposit) error
	GetUserDeposit(userID string, startDate, endDate time.Time) ([]*entity.Deposit, error)
}
