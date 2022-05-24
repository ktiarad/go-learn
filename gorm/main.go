package main

import (
	"fmt"
	"gorm/database"
	"gorm/repository"
	"strings"
)

func main() {
	db := database.StartDb()
	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)

	// ########## USER ##########

	// ADD NEW USER
	// user := models.User{
	// 	Email: "shino@konoha.com",
	// }

	// err := userRepo.CreateUser(&user)
	// if err != nil {
	// 	fmt.Println("error:", err.Error())
	// 	return
	// }

	// fmt.Println("Created success")

	// GET USER BY ID
	// emp, err := userRepo.GetUserById(2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// emp.Print()

	// UPDATE USER
	// user, err := userRepo.UpdateUserById(3, "hinata@konoha.com")
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("Email has been changed :", user.Email)

	// DELETE USER
	// var id uint = 4
	// err := userRepo.DeleteUserById(id)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("A user has been deleted with id:", id)

	// GET ALL USERS
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

	// ########## PRODUCT ##########
	// ADD NEW PRODUCT
	// product := models.Product{
	// 	Name:   "Baju",
	// 	Brand:  "XYZ",
	// 	UserID: 5,
	// }

	// err := productRepo.CreateProduct(&product)
	// if err != nil {
	// 	fmt.Println("error:", err.Error())
	// 	return
	// }
	// fmt.Println("Created success")

	// GET PRODUCT BY ID
	// product, err := productRepo.GetProductById(3)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// product.Print()

	// UPDATE PRODUCT
	// product, err := productRepo.UpdateProductById(3, "Kaos")
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("Product has been changed :", product.Name)

	// DELETE PRODUCT
	// var id uint = 4
	// err := productRepo.DeleteProductById(4)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("A product has been deleted with id:", id)

	// GET ALL PRODUCTS
	products, err := productRepo.GetAllProducts()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for k, prod := range *products {
		fmt.Println("Product :", k+1)
		prod.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

}
