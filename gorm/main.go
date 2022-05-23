package main

import (
	"fmt"
	"gorm/database"
	"gorm/models"
	"gorm/repository"
	"strings"
)

func main() {
	db := database.StartDb()

	user := models.User{
		Email: "naruto@konoha.com",
	}

	userRepo := repository.NewUserRepo(db)
	err := userRepo.CreateUser(&user)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	fmt.Println("Created success")

	employees, err := userRepo.GetAllUsers()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for k, emp := range *employees {
		fmt.Println("User :", k+1)
		emp.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	emp, err := userRepo.GetUserById(3)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	emp.Print()
}
