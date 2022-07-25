package business

import "Taller-LCC/Clase1/domain"

type ProductBusiness interface {
	GetAllProducts() ([]domain.Product, error)
	GetAllProductsById(int) ([]domain.Product, error)
	UpdateProductById(int) (domain.Product, error)
}
