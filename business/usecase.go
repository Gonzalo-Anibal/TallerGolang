package business

import (
	"Taller-LCC/Clase1/domain"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/model"
)

type ProductBusiness interface {
	GetAllProducts() ([]domain.Product, error)
	NewProduct(product model.Product) ([]domain.Product, error)
}
