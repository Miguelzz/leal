package product_repo

import (
	"loyalty/internal/core/domain"
	"loyalty/internal/repositories/mocks"
)

type productTable struct {
}

func New() *productTable {
	return &productTable{}
}

func (repo *productTable) All() ([]domain.Product, error) {
	return mocks.MockProducts, nil
}

func (repo *productTable) Get(id int) (domain.Product, error) {
	return mocks.MockProducts[id], nil
}
