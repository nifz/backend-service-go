package routes

import (
	"backend-service-go/controllers"
	"backend-service-go/repositories"
	"backend-service-go/usecases"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	productRepository := repositories.NewProductRepository(db)
	productUsecase := usecases.NewProductUsecase(productRepository)
	productController := controllers.NewProductController(productUsecase)

	cartRepository := repositories.NewCartRepository(db)
	cartUsecase := usecases.NewCartUsecase(cartRepository, productRepository)
	cartController := controllers.NewCartController(cartUsecase)

	e.GET("/products", productController.GetAllProduct)
	e.GET("/product/:id", productController.GetProductByCode)
	e.POST("/product", productController.CreateProduct)
	e.PATCH("/product/:id", productController.UpdateProduct)
	e.DELETE("/product/:id", productController.DeleteProduct)

	e.GET("/carts", cartController.GetAllCart)
	e.POST("/cart", cartController.CreateCart)
	e.PATCH("/cart/:id", cartController.UpdateCart)
	e.DELETE("/cart/:id", cartController.DeleteCart)

}
