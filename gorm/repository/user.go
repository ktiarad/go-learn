package repository

import (
	"gorm/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(*models.User) error
	GetAllUsers() (*[]models.User, error)
	GetUserById(id uint) (*models.User, error)
	UpdateUserById(id uint, email string) (*models.User, error)
	DeleteUserById(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(request *models.User) error {
	err := r.db.Create(request).Error
	return err

}
func (r *userRepo) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return &users, err
}
func (r *userRepo) GetUserById(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "id=?", id).Error
	return &user, err
}

func (r *userRepo) UpdateUserById(id uint, email string) (*models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", id).Updates(models.User{Email: email}).Error
	return &user, err
}

func (r *userRepo) DeleteUserById(id uint) error {
	var user models.User

	err := r.db.Where("id = ?", id).Delete(&user).Error

	return err
}
