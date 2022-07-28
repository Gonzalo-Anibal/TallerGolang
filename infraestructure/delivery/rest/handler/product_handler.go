package handler

import (
	"Taller-LCC/Clase1/business"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/model"
	"github.com/labstack/echo"
	"net/http"
)

type productHandler struct {
	useCaseProduct business.ProductBusiness
}

func NewProductHandler(e *echo.Echo, productBusiness business.ProductBusiness) *productHandler {
	p := &productHandler{
		useCaseProduct: productBusiness,
	}
	e.GET("/products", p.getAllProductsMock)
	e.POST("/product", p.postProduct)
	return p
}

func (p *productHandler) getAllProductsMock(c echo.Context) error {
	result, err := p.useCaseProduct.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	return c.JSON(http.StatusOK, result)
}

func (p *productHandler) postProduct(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	products, err := p.useCaseProduct.NewProduct(*product)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": err})
	}
	return c.JSON(http.StatusOK, products)
}
