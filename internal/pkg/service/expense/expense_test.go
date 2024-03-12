package expense

import (
	"errors"
	"testing"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/deBeloper-code/goFinance/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type expenseServiceSuite struct {
	suite.Suite
	//Dependencies
	repo *mocks.ExpenseRepository
	// Service to test
	service ports.ExpenseService
}

// Running Before each test
func (suite *expenseServiceSuite) SetupTest() {
	suite.repo = new(mocks.ExpenseRepository)
	suite.service = NewService(suite.repo)
}

// Running After each test
func (suite *expenseServiceSuite) TearDownTest() {
	suite.SetupTest()
}

func (suite *expenseServiceSuite) TestAddNegative() {
	// Mocking result returned by our repository
	suite.repo.On("AddExpense", mock.Anything).Return(errors.New("Something was wrong"))

	err := suite.service.Add(&entity.Expense{
		UserId:   "1",
		Category: "Test",
	})
	suite.Error(err)
}
func (suite *expenseServiceSuite) TestGetUser() {
	expectedExpenses := []*entity.Expense{
		{
			Category: "Dummy test expense",
			Name:     "category expense 1",
		},
		{
			Category: "Dummy test expense",
			Name:     "category expense 2",
		},
	}

	// Mocking result returned by our repository
	suite.repo.On("GetUserExpense",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("time.Time"),
		mock.AnythingOfType("time.Time"),
	).Return(expectedExpenses, nil)

	expenses, err := suite.service.GetUserExpense("21", "2024-01-08", "2024-01-08")
	suite.Len(expenses, 2)
	suite.ElementsMatch(expenses, expectedExpenses)
	suite.NoError(err)
}

func TestExpectService(t *testing.T) {
	suite.Run(t, new(expenseServiceSuite))
}
