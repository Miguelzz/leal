package points_srv

import (
	"errors"
	"fmt"
	"loyalty/internal/core/domain"
	"loyalty/internal/core/ports"
)

type service struct {
	pointsEvents ports.PointsEvent

	pointsRepository   ports.PointsRepository
	productsRepository ports.ProductRepository
}

func New(pointsEvents ports.PointsEvent, pointsRepository ports.PointsRepository, productsRepository ports.ProductRepository) *service {
	return &service{
		pointsEvents:       pointsEvents,
		pointsRepository:   pointsRepository,
		productsRepository: productsRepository,
	}
}

func (srv *service) Get(id int) (domain.Points, error) {
	points, err := srv.pointsRepository.Get(id)
	if err != nil {
		return domain.Points{}, errors.New("get points from repository has failed")
	}

	return points, nil
}

func (srv *service) Redeem(id int, idProduct int) (domain.Points, error) {

	points, err := srv.pointsRepository.Get(id)
	product, err := srv.productsRepository.Get(idProduct)
	points.Points -= product.Price/1000 + 5
	res, err := srv.pointsEvents.Redeem(points)
	newPoints, err := srv.pointsRepository.Update(points)
	fmt.Println(res, points)

	return newPoints, err
}

func (srv *service) Buy(id int, idProduct int) (domain.Points, error) {
	points, err := srv.pointsRepository.Get(id)
	product, err := srv.productsRepository.Get(idProduct)
	points.Points += product.Price/5000 + 5
	res, err := srv.pointsEvents.Buy(points)
	newPoints, err := srv.pointsRepository.Update(points)
	fmt.Println(res, points)

	return newPoints, err
}
