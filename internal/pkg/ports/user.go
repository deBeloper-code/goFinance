package ports

type UserRepository interface {
	Create(value interface{}) error
	First(dest interface{}, conds ...interface{}) error
}
