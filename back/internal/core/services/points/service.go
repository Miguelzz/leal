package points_srv

import (
	"errors"
	"loyalty/internal/core/domain"
	"loyalty/internal/core/ports"
)

type service struct {
	pointsRepository   ports.PointsRepository
	productsRepository ports.ProductRepository
}

func New(pointsRepository ports.PointsRepository, productsRepository ports.ProductRepository) *service {
	return &service{
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
	newPoints, err := srv.pointsRepository.Update(points)

	if err != nil {
		return domain.Points{}, errors.New("get points from repository has failed")
	}

	return newPoints, nil
}

func (srv *service) Buy(id int, idProduct int) (domain.Points, error) {
	points, err := srv.pointsRepository.Get(id)
	product, err := srv.productsRepository.Get(idProduct)
	points.Points += product.Price/5000 + 5
	newPoints, err := srv.pointsRepository.Update(points)

	if err != nil {
		return domain.Points{}, errors.New("get points from repository has failed")
	}

	return newPoints, nil
}

func (srv *service) Update(idProduct int, points domain.Points) (domain.Points, error) {

	product, err := srv.productsRepository.Get(idProduct)

	if err != nil {
		return domain.Points{}, errors.New("get product from repository has failed")
	}

	newPoints, err := srv.pointsRepository.Update(domain.Points{
		ID:     points.ID,
		Points: product.Price/5000 + 5,
	})

	if err != nil {
		return domain.Points{}, errors.New("update points repository has failed")
	}

	return newPoints, nil
}
