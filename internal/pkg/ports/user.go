package ports

import "github.com/deBeloper-code/goFinance/internal/pkg/entity"

type UserRepository interface {
	Create(value interface{}) error
	First(dest interface{}, conds ...interface{}) error
}
type UserService interface {
	Create(user *entity.User) error
	Login(credentials *entity.DefaultCredentials) (string, error)
}
