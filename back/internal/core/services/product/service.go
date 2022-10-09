package product_srv

import (
	"errors"
	"loyalty/internal/core/domain"
	"loyalty/internal/core/ports"
)

type service struct {
	productsRepository ports.ProductRepository
}

func New(productsRepository ports.ProductRepository) *service {
	return &service{
		productsRepository: productsRepository,
	}
}

func (srv *service) Get(id int) (domain.Product, error) {
	product, err := srv.productsRepository.Get(id)
	if err != nil {
		return domain.Product{}, errors.New("get product from repository has failed")
	}

	return product, nil
}

func (srv *service) All() ([]domain.Product, error) {
	products, err := srv.productsRepository.All()

	if err != nil {
		return []domain.Product{}, errors.New("get product from repository has failed")
	}

	return products, nil
}
