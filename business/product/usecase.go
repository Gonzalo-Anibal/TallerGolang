package product

import (
	"Taller-LCC/Clase1/domain"
	"Taller-LCC/Clase1/infraestructure/repository"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/model"
)

type productUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) *productUseCase {
	return &productUseCase{
		repository: productRepository,
	}
}

func (p *productUseCase) GetAllProducts() ([]domain.Product, error) {
	return p.repository.GetAll()
}

func (p *productUseCase) NewProduct(product model.Product) ([]domain.Product, error) {
	return p.repository.NewProduct(product)
}
