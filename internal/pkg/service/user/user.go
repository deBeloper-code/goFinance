package user

import (
	"net/mail"

	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo ports.UserRepository
}

func NewService(repo ports.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(user *entity.User) error {
	// 1. Valid email address
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}
	// 2. Hash password
	user.Password = hashAndSalt(user.Password)
	// 3 Save into DB
	return s.repo.Create(user)
}

func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}
	return string(hash)
}

func (s *service) Login(credentials *entity.DefaultCredentials) (string, error) {
	user := &entity.User{}
	// 1. Looking for if email exists
	if err := s.repo.First(user, "email = ?", credentials.Email); err != nil {
		return "", err
	}
	// 2. Trying match password
	if err := tryMatchPassword(user.Password, credentials.Password); err != nil {
		return "", err
	}
	// 3. Create Session token JWT

}

func tryMatchPassword(userPassword, credentialsPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(credentialsPassword))
	if err != nil {
		return err
	}
	return nil
}
