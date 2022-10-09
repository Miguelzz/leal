package domain

type User struct {
	ID       int `json:"id"`
	IdPoints int `json:"idPoints"`
}

type Points struct {
	ID     int `json:"id"`
	Points int `json:"points"`
}

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
