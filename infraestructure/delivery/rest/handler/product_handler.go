package handler

type productHandler interface {
	useCase business.ProductBusiness
}

func NewProductHandler() {}
