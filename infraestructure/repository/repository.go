package repository

import "Taller-LCC/Clase1/domain"

type ProductRepository interface {
	GetAll() ([]domain.Product, error)
}
