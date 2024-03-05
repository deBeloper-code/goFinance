package ports

import (
	"time"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
)

type ExpenseRepository interface {
	AddExpense(expense *entity.Expense) error
	GetUserExpense(userID string, startDate, endDate time.Time) ([]*entity.Expense, error)
}
