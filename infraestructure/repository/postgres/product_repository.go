package postgres

import (
	"Taller-LCC/Clase1/domain"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/mapper"
	"Taller-LCC/Clase1/infraestructure/repository/postgres/model"
	"errors"
)

type productRepository struct {
	client string
}

func NewProductRepository(connection string) *productRepository {
	return &productRepository{
		client: connection,
	}
}

func (r *productRepository) GetAll() ([]domain.Product, error) {
	productsDB, err := resultAllProductsFromDB()
	if err != nil {
		return nil, errors.New("Error en obtener Productos Mock")
	}
	products := mapper.ModelsToDomainsProduct(productsDB)
	return products, nil
}

func (r *productRepository) NewProduct(product model.Product) ([]domain.Product, error) {
	productsDB, err := resultAllProductsFromDB()
	if err != nil {
		return nil, errors.New("Error en obtener Productos Mock")
	}
	productsDB = append(productsDB, product)
	products := mapper.ModelsToDomainsProduct(productsDB)

	return products, nil
}

var productsMock = []model.Product{
	{
		UUID:        1,
		Id:          1,
		Brand:       "Apple",
		Description: "iPhone 12 Pro 128 Dorado",
		Image:       "https://images.lider.cl/wmtcl?source=url[file:/productos/1278763a.jpg]&sink",
		Price:       989990,
	},
	{
		UUID:        2,
		Id:          2,
		Brand:       "Microsoft",
		Description: "Consola Xbox Series S",
		Image:       "https://images.lider.cl/wmtcl?source=url[file:/productos/1094365a.jpg]&sink",
		Price:       329990,
	},
}

func resultAllProductsFromDB() ([]model.Product, error) {
	return productsMock, nil
}
