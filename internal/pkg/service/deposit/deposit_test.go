package deposit

import (
	"errors"
	"testing"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/deBeloper-code/goFinance/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type depositServiceSuite struct {
	suite.Suite
	//Dependencies
	repo *mocks.DepositRepository
	// Service to test
	service ports.DepositService
}

// Running Before each test
func (suite *depositServiceSuite) SetupTest() {
	suite.repo = new(mocks.DepositRepository)
	suite.service = NewService(suite.repo)
}

// Running After each test
func (suite *depositServiceSuite) TearDownTest() {
	suite.SetupTest()
}

func TestExpectService(t *testing.T) {
	suite.Run(t, new(depositServiceSuite))
}

func (suite *depositServiceSuite) TestAddNegative() {
	// Mocking result returned by our repository
	suite.repo.On("AddDeposit", mock.Anything).Return(errors.New("Something was wrong"))

	err := suite.service.Add(&entity.Deposit{
		UserId:   "1",
		Category: "Test",
	})
	suite.Error(err)
}

func (suite *depositServiceSuite) TestGetUser() {
	expectedDeposits := []*entity.Deposit{
		{
			Category: "Dummy test deposit",
			Name:     "category expense 1",
		},
		{
			Category: "Dummy test deposit",
			Name:     "category deposit 2",
		},
	}

	// Mocking result returned by our repository
	suite.repo.On("GetUserDeposit",
		mock.AnythingOfType("string"),
		mock.AnythingOfType("time.Time"),
		mock.AnythingOfType("time.Time"),
	).Return(expectedDeposits, nil)

	deposits, err := suite.service.GetUserDeposit("21", "2024-01-08", "2024-01-08")
	suite.Len(deposits, 2)
	suite.ElementsMatch(deposits, expectedDeposits)
	suite.NoError(err)
}
