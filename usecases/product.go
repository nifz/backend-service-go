package usecases

import (
	"backend-service-go/dtos"
	"backend-service-go/models"
	"backend-service-go/repositories"
)

type ProductUsecase interface {
	CreateProduct(productInput dtos.ProductInput) (dtos.ProductResponse, error)
	GetAllProduct() ([]dtos.ProductResponse, error)
	GetProductByCode(productCode string) (dtos.ProductResponse, error)
	UpdateProduct(productInput dtos.ProductInput, productCode string) (dtos.ProductResponse, error)
	DeleteProduct(productCode string) error
}

type productUsecase struct {
	repository repositories.ProductRepository
}

func NewProductUsecase(r repositories.ProductRepository) *productUsecase {
	return &productUsecase{r}
}

func (uc *productUsecase) CreateProduct(productInput dtos.ProductInput) (dtos.ProductResponse, error) {
	var productResponse dtos.ProductResponse
	products := models.Product{
		Code:        productInput.CodeProduct,
		Name:        productInput.Name,
		Image:       productInput.Image,
		Description: productInput.Description,
		Price:       productInput.Price,
		Quantity:    productInput.Quantity,
	}
	product, err := uc.repository.CreateProduct(products)
	if err != nil {
		return productResponse, err
	}
	productResponse = dtos.ProductResponse{
		CodeProduct: product.Code,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
	}
	return productResponse, err
}

func (uc *productUsecase) GetAllProduct() ([]dtos.ProductResponse, error) {
	var productResponses []dtos.ProductResponse

	products, err := uc.repository.GetAllProduct()
	if err != nil {
		return productResponses, err
	}
	for _, product := range products {
		productResponse := dtos.ProductResponse{
			CodeProduct: product.Code,
			Name:        product.Name,
			Image:       product.Image,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			CreatedAt:   &product.CreatedAt,
			UpdatedAt:   &product.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}

func (uc *productUsecase) GetProductByCode(productCode string) (dtos.ProductResponse, error) {
	var productResponse dtos.ProductResponse

	product, err := uc.repository.GetProductByCode(productCode)
	if err != nil {
		return productResponse, err
	}
	productResponse = dtos.ProductResponse{
		CodeProduct: product.Code,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
	}
	return productResponse, err
}

func (uc *productUsecase) UpdateProduct(productInput dtos.ProductInput, productCode string) (dtos.ProductResponse, error) {
	var productResponse dtos.ProductResponse

	product, err := uc.repository.GetProductByCode(productCode)
	if err != nil {
		return productResponse, err
	}

	product.Name = productInput.Name
	product.Image = productInput.Image
	product.Description = productInput.Description
	product.Price = productInput.Price
	product.Quantity = productInput.Quantity

	product, err = uc.repository.UpdateProduct(product)
	if err != nil {
		return productResponse, err
	}

	productResponse = dtos.ProductResponse{
		CodeProduct: product.Code,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   &product.CreatedAt,
		UpdatedAt:   &product.UpdatedAt,
	}

	return productResponse, nil
}

func (uc *productUsecase) DeleteProduct(productCode string) error {
	product, err := uc.repository.GetProductByCode(productCode)

	if err != nil {
		return err
	}

	if productCode == product.Code {
		uc.repository.DeleteProduct(product)
	} else {
		return err
	}

	return nil
}
