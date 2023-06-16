package repositories

import (
	"backend-service-go/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	GetAllCart() ([]models.Cart, error)
	GetProductByCode(productCode string) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) error
}

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (repo *cartRepository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := repo.DB.Create(&cart).Error
	if err != nil {
		return models.Cart{}, err
	}
	err = repo.DB.Preload("Product").Where("code_product = ?", cart.CodeProduct).First(&cart).Error
	if err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

func (repo *cartRepository) GetAllCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := repo.DB.Preload("Product").Find(&cart).Error
	return cart, err
}

func (repo *cartRepository) GetProductByCode(productCode string) (models.Cart, error) {
	var cart models.Cart
	err := repo.DB.Preload("Product").Where("code_product = ?", productCode).First(&cart).Error
	return cart, err
}

func (repo *cartRepository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := repo.DB.Preload("Product").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart).Where("id = ?", cart.ID).First(&cart).Error
	return cart, err
}

func (repo *cartRepository) DeleteCart(cart models.Cart) error {
	err := repo.DB.Delete(&cart).Error
	return err
}
