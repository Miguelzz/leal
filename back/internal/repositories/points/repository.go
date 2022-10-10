package points_repo

import (
	"fmt"
	"loyalty/internal/core/domain"
)

var users = []domain.User{
	{
		ID:       1,
		IdPoints: 1,
	},
}

var points = []domain.Points{
	{
		ID:     1,
		Points: 0,
	},
}

type PointsTable struct {
}

func New() *PointsTable {
	return &PointsTable{}
}

func (repo *PointsTable) Get(id int) (domain.Points, error) {
	fmt.Println(points[0])
	return points[id-1], nil
}

func (repo *PointsTable) Update(_points domain.Points) (domain.Points, error) {
	points[_points.ID-1] = _points
	return _points, nil
}
