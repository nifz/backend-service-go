package usecases

import (
	"backend-service-go/dtos"
	"backend-service-go/models"
	"backend-service-go/repositories"
)

type CartUsecase interface {
	CreateCart(cartInput dtos.CartInput) (dtos.CartResponse, error)
	GetAllCart() ([]dtos.CartResponse, error)
	UpdateCart(cartInput dtos.CartInput, productCode string) (dtos.CartResponse, error)
	DeleteCart(cartCode string) error
}

type cartUsecase struct {
	cartRepository    repositories.CartRepository
	productRepository repositories.ProductRepository
}

func NewCartUsecase(cartRepository repositories.CartRepository, productRepository repositories.ProductRepository) *cartUsecase {
	return &cartUsecase{cartRepository, productRepository}
}

func (uc *cartUsecase) CreateCart(cartInput dtos.CartInput) (dtos.CartResponse, error) {
	var cartResponse dtos.CartResponse
	cart := models.Cart{
		CodeProduct: cartInput.CodeProduct,
		Quantity:    cartInput.Quantity,
	}
	cart, err := uc.cartRepository.CreateCart(cart)
	if err != nil {
		return dtos.CartResponse{}, err
	}
	product, err := uc.productRepository.GetProductByCode(cart.CodeProduct)
	if err != nil {
		return cartResponse, err
	}
	product.Quantity -= cartInput.Quantity
	product, err = uc.productRepository.UpdateProduct(product)
	if err != nil {
		return cartResponse, err
	}
	product, err = uc.productRepository.GetProductByCode(cart.CodeProduct)
	if err != nil {
		return dtos.CartResponse{}, err
	}

	cartResponse = dtos.CartResponse{
		ID: int(cart.ID),
		Product: dtos.ProductResponse{
			CodeProduct: product.Code,
			Name:        product.Name,
			Image:       product.Image,
			Description: product.Description,
			Price:       product.Price,
		},
		Quantity:  cart.Quantity,
		CreatedAt: &cart.CreatedAt,
		UpdatedAt: &cart.UpdatedAt,
	}

	return cartResponse, err
}

func (uc *cartUsecase) GetAllCart() ([]dtos.CartResponse, error) {
	var cartResponses []dtos.CartResponse

	carts, err := uc.cartRepository.GetAllCart()
	if err != nil {
		return cartResponses, err
	}

	for _, cart := range carts {
		product, err := uc.productRepository.GetProductByCode(cart.CodeProduct)
		if err != nil {
			return cartResponses, err
		}
		cartResponse := dtos.CartResponse{
			ID: int(cart.ID),
			Product: dtos.ProductResponse{
				CodeProduct: product.Code,
				Name:        product.Name,
				Image:       product.Image,
				Description: product.Description,
				Price:       product.Price,
			},
			Quantity:  cart.Quantity,
			CreatedAt: &cart.CreatedAt,
			UpdatedAt: &cart.UpdatedAt,
		}
		cartResponses = append(cartResponses, cartResponse)
	}
	if err != nil {
		return cartResponses, err
	}

	return cartResponses, nil
}

func (uc *cartUsecase) UpdateCart(cartInput dtos.CartInput, productCode string) (dtos.CartResponse, error) {
	var cartResponse dtos.CartResponse
	cart, err := uc.cartRepository.GetProductByCode(productCode)
	if err != nil {
		return cartResponse, err
	}
	productQty := 0
	if cartInput.Quantity > cart.Quantity {
		productQty = cartInput.Quantity - cart.Quantity
	} else if cartInput.Quantity == cart.Quantity {
		productQty = cart.Quantity
	} else {
		productQty = cart.Quantity - cartInput.Quantity
	}
	cart.Quantity = cartInput.Quantity

	cart, err = uc.cartRepository.UpdateCart(cart)

	product, err := uc.productRepository.GetProductByCode(cart.CodeProduct)
	if err != nil {
		return cartResponse, err
	}
	product.Quantity = productQty
	product, err = uc.productRepository.UpdateProduct(product)
	if err != nil {
		return cartResponse, err
	}
	cartResponse = dtos.CartResponse{
		ID: int(cart.ID),
		Product: dtos.ProductResponse{
			CodeProduct: product.Code,
			Name:        product.Name,
			Image:       product.Image,
			Description: product.Description,
			Price:       product.Price,
		},
		Quantity:  cart.Quantity,
		CreatedAt: &cart.CreatedAt,
		UpdatedAt: &cart.UpdatedAt,
	}
	if err != nil {
		return cartResponse, err
	}

	return cartResponse, nil
}

func (uc *cartUsecase) DeleteCart(productCode string) error {
	cart, err := uc.cartRepository.GetProductByCode(productCode)

	if err != nil {
		return err
	}

	if productCode == cart.CodeProduct {
		uc.cartRepository.DeleteCart(cart)
	} else {
		return err
	}

	return nil
}
