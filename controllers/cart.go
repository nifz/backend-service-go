package controllers

import (
	"backend-service-go/dtos"
	"backend-service-go/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	cartUsecase usecases.CartUsecase
}

func NewCartController(cartUsecase usecases.CartUsecase) CartController {
	return CartController{cartUsecase}
}

func (ctrl *CartController) CreateCart(c echo.Context) error {
	var cart dtos.CartInput

	err := c.Bind(&cart)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to input cart",
			})
	}

	carts, err := ctrl.cartUsecase.CreateCart(cart)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to input cart",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully submitted cart",
			"data":    carts,
		})
}

func (ctrl *CartController) GetAllCart(c echo.Context) error {
	carts, err := ctrl.cartUsecase.GetAllCart()
	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Failed to get all cart",
			},
		)
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully retrieved all cart",
			"data":    carts,
		})
}

func (ctrl *CartController) UpdateCart(c echo.Context) error {
	var cart dtos.CartInput

	byID := c.Param("id")

	err := c.Bind(&cart)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cannot input the data",
			})
	}

	carts, err := ctrl.cartUsecase.UpdateCart(cart, byID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cart ID cannot be found",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully updated cart",
			"data":    carts,
		})
}

func (ctrl *CartController) DeleteCart(c echo.Context) error {
	byID := c.Param("id")

	err := ctrl.cartUsecase.DeleteCart(byID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest, echo.Map{
				"message": "Cart ID cannot be found",
			})
	}

	return c.JSON(
		http.StatusOK, echo.Map{
			"message": "Successfully deleted cart",
		})
}
