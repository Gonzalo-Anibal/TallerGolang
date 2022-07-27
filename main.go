package main

import (
	"Taller-LCC/Clase1/business/product"
	"Taller-LCC/Clase1/infraestructure/delivery/rest/handler"
	"Taller-LCC/Clase1/infraestructure/repository/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	connection := "Esta es una conexion Mock"

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.HideBanner = true

	productRepository := postgres.NewProductRepository(connection)
	productUseCase := product.NewProductUseCase(productRepository)

	handler.NewProductHandler(e, productUseCase)

	log.Printf("Starting server")
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(server))
}
