package mapper

import (
	"Taller-LCC/Clase1/domain"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/model"
)

func ModelsToDomainsProduct(productsDB []model.Product) []domain.Product {
	products := make([]domain.Product, 0)
	for _, productsDB := range productsDB {
		products = append(products, ModelToDomainProduct(productsDB))
	}
	return products
}

func ModelToDomainProduct(productDB model.Product) domain.Product {
	return domain.Product{
		Id:          productDB.Id,
		Brand:       productDB.Brand,
		Description: productDB.Description,
		Image:       productDB.Image,
		Price:       productDB.Price,
	}
}
