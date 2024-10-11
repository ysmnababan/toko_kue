package main

import (
	"log"
	"os"
	"toko_kue/config"
	"toko_kue/handler"
	"toko_kue/models"
	"toko_kue/repository"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}


func main() {
	db := config.Connect()
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	e := echo.New()

	categoryRepo := &repository.Repo{DB: db}
	categoryHandler := &handler.CategoryHandler{CR: categoryRepo}

	productRepo := &repository.Repo{DB: db}
	productHandler := &handler.ProductHandler{PR: productRepo}
	// Configure CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500"}, // Replace with your frontend URL
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/categories", categoryHandler.GetAllCategory)
	e.GET("/categories/:id", categoryHandler.GetById)
	e.POST("/categories", categoryHandler.AddCategory)
	e.PUT("/categories/:id", categoryHandler.UpdateCategory)
	e.DELETE("/categories/:id", categoryHandler.DeleteCategory)

	e.GET("/products", productHandler.GetAllProduct)
	e.GET("/products/:id", productHandler.GetProductById)
	e.POST("/products", productHandler.AddProduct)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)

	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}
