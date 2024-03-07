package main

import (
	"fmt"

	"github.com/deBeloper-code/goFinance/internal/infra/repositories/postgres"
	"github.com/deBeloper-code/goFinance/internal/pkg/entity"
	"github.com/deBeloper-code/goFinance/internal/pkg/service/user"
)

func main() {
	userRepository := postgres.NewClient()
	userService := user.NewService(userRepository)
	// err := userService.Create(&entity.User{
	// 	ID:       uuid.NewString(),
	// 	Name:     "bryan",
	// 	LastName: "sanchez",
	// 	Email:    "bry@hola.com",
	// 	Password: "hola123456",
	// })
	// fmt.Println("My error:", err)
	token, err := userService.Login(&entity.DefaultCredentials{
		Email:    "bry@hola.com",
		Password: "hola123456",
	})
	fmt.Println(token, err)
}
