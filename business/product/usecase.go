package product

import (
	"Taller-LCC/Clase1/domain"
	"Taller-LCC/Clase1/infraestructure/repository"
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
