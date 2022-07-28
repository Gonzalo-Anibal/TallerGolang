package repository

import (
	"Taller-LCC/Clase1/domain"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/model"
)

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
	NewProduct(product model.Product) ([]domain.Product, error)
}
