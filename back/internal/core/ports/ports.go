package ports

import "loyalty/internal/core/domain"

type UserRepository interface {
	Get(id int) (domain.User, error)
}

type PointsRepository interface {
	Get(id int) (domain.Points, error)
}

type ProductRepository interface {
	All() ([]domain.Product, error)
	Get(id int) (domain.Product, error)
}

type UserService interface {
	Get(id int) (domain.User, error)
}

type PointsService interface {
	Get(id int) (domain.Points, error)
	Buy(id int, idProduct int) error
	Redeem(id int, idProduct int) error
}

type ProductService interface {
	All() ([]domain.Product, error)
	Get(id int) (domain.Product, error)
}

type PointsEvent interface {
	Buy(points domain.Points) (int, error)
	Redeem(points domain.Points) (int, error)
}
