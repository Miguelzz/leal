package points_repo

import (
	"loyalty/internal/core/domain"
	"loyalty/internal/repositories/mocks"
)

type pointsTable struct {
}

func New() *pointsTable {
	return &pointsTable{}
}

func (repo *pointsTable) Get(id int) (domain.Points, error) {
	return mocks.MockReadPoints[id], nil
}

func (repo *pointsTable) Update(_points domain.Points) (domain.Points, error) {
	mocks.MockWriterPoints[_points.ID] = _points
	return _points, nil
}
