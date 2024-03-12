package user

import (
	"testing"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/deBeloper-code/goFinance/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userServiceSuite struct {
	suite.Suite
	//Dependencies
	repo *mocks.UserRepository
	// Service to test
	service ports.UserService
}

// Running Before each test
func (suite *userServiceSuite) SetupTest() {
	suite.repo = new(mocks.UserRepository)
	suite.service = NewService(suite.repo)
}

// Running After each test
func (suite *userServiceSuite) TearDownTest() {
	suite.SetupTest()
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}

func (suite *userServiceSuite) TestCreate() {
	suite.repo.On("Create", mock.Anything).Return(nil)

	err := suite.service.Create(&entity.User{
		Email:    "josue140596@gmail.com",
		Password: "123pass",
	})

	suite.NoError(err)
}
