package repositories

import (
	"backend-service-go/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetAllProduct() ([]models.Product, error)
	GetProductByCode(productCode string) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) error
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (repo *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	err := repo.DB.Create(&product).Error
	return product, err
}

func (repo *productRepository) GetAllProduct() ([]models.Product, error) {
	var product []models.Product
	err := repo.DB.Find(&product).Error
	return product, err
}

func (repo *productRepository) GetProductByCode(productCode string) (models.Product, error) {
	var product models.Product
	err := repo.DB.Where("code = ?", productCode).First(&product).Error
	return product, err
}

func (repo *productRepository) UpdateProduct(product models.Product) (models.Product, error) {
	err := repo.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&product).Where("code = ?", product.Code).First(&product).Error
	return product, err
}

func (repo *productRepository) DeleteProduct(product models.Product) error {
	err := repo.DB.Delete(&product).Error
	return err
}
