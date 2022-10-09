package ports

import "loyalty/internal/core/domain"

type UserRepository interface {
	Get(id int) (domain.User, error)
	Update(domain.User) (domain.User, error)
}

type PointsRepository interface {
	Get(id int) (domain.Points, error)
	Update(domain.Points) (domain.Points, error)
}

type ProductRepository interface {
	All() ([]domain.Product, error)
	Get(id int) (domain.Product, error)
}

type UserService interface {
	Get(id int) (domain.User, error)
	Update(domain.User) (domain.User, error)
}

type PointsService interface {
	Get(id int) (domain.Points, error)
	Update(idProduct int, points domain.Points) (domain.Points, error)
}

type ProductService interface {
	All() ([]domain.Product, error)
	Get(id int) (domain.Product, error)
}
