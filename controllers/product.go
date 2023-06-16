package controllers

import (
	"backend-service-go/dtos"
	"backend-service-go/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUsecase usecases.ProductUsecase
}

func NewProductController(productUsecase usecases.ProductUsecase) ProductController {
	return ProductController{productUsecase}
}

func (ctrl *ProductController) CreateProduct(c echo.Context) error {
	var product dtos.ProductInput

	err := c.Bind(&product)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to input product",
			})
	}

	products, err := ctrl.productUsecase.CreateProduct(product)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to input product",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully submitted product",
			"data":    products,
		})
}

func (ctrl *ProductController) GetAllProduct(c echo.Context) error {
	products, err := ctrl.productUsecase.GetAllProduct()

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get product",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all product",
			"data":    products,
		})
}

func (ctrl *ProductController) GetProductByCode(c echo.Context) error {
	byID := c.Param("id")

	product, err := ctrl.productUsecase.GetProductByCode(byID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get product",
			})
	}

	if len(product.CodeProduct) < 1 {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get data, id is invalid",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved product by code",
			"data":    product,
		})
}

func (ctrl *ProductController) UpdateProduct(c echo.Context) error {
	var product dtos.ProductInput

	byID := c.Param("id")

	err := c.Bind(&product)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cannot input the data",
			})
	}

	products, err := ctrl.productUsecase.UpdateProduct(product, byID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Product ID cannot be found",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully updated product",
			"data":    products,
		})
}

func (ctrl *ProductController) DeleteProduct(c echo.Context) error {
	byID := c.Param("id")

	err := ctrl.productUsecase.DeleteProduct(byID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Product ID cannot be found",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully deleted product",
		})
}
