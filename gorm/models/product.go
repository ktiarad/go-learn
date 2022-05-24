package models

import (
	"fmt"
	"time"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (u *Product) BeforeCreate(db *gorm.DB) (err error) {
// 	fmt.Println("Before insert to table users")
// 	if len(u.Email) < 10 {
// 		err = fmt.Errorf("your email is too short")
// 	}
// 	// err = nil
// 	return err
// }

func (u *Product) Print() {
	fmt.Println("ID :", u.ID)
	fmt.Println("Name :", u.Name)
	fmt.Println("Brand :", u.Brand)
	fmt.Println("UserID :", u.UserID)
}
