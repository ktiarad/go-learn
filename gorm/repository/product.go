package repository

import (
	"gorm/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(*models.Product) error
	GetAllProducts() (*[]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	UpdateProductById(id uint, email string) (*models.Product, error)
	DeleteProductById(id uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) CreateProduct(request *models.Product) error {
	err := r.db.Create(request).Error
	return err
}

func (r *productRepo) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return &products, err
}
func (r *productRepo) GetProductById(id uint) (*models.Product, error) {
	var product models.Product

	err := r.db.First(&product, "id=?", id).Error
	return &product, err
}

func (r *productRepo) UpdateProductById(id uint, name string) (*models.Product, error) {
	var products models.Product

	err := r.db.Where("id = ?", id).Updates(models.Product{Name: name}).Error
	return &products, err
}

func (r *productRepo) DeleteProductById(id uint) error {
	var product models.Product

	err := r.db.Where("id = ?", id).Delete(&product).Error

	return err
}
