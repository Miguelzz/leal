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
	return points[0], nil
}

func (repo *PointsTable) Update(_points domain.Points) (domain.Points, error) {
	points[0] = _points
	return _points, nil
}

// type PointsBalance struct {
// 	Points  int `json:"points"`
// 	Balance int `json:"balance"`
// }

// var points = 0
// var balance = 200_000

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	products := prepareResponse()
// 	jsonResponse, err := json.Marshal(products)
// 	if err != nil {
// 		return
// 	}
// 	w.Write(jsonResponse)
// }

// func GetState(w http.ResponseWriter, r *http.Request) {
// 	jsonResponse, err := json.Marshal(PointsBalance{
// 		Points:  points,
// 		Balance: balance,
// 	})

// 	if err != nil {
// 		return
// 	}
// 	w.Write(jsonResponse)
// }

// func PostBuy(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(mux.Vars(r)["id"])
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(id)
// 	price := prepareResponse()[id-1].Price

// 	points += price/1000 + 5
// 	balance = balance - price

// 	jsonResponse, err := json.Marshal(PointsBalance{
// 		Points:  points,
// 		Balance: balance,
// 	})

// 	if err != nil {
// 		return
// 	}
// 	w.Write(jsonResponse)
// }

// func PostRedeem(w http.ResponseWriter, r *http.Request) {

// 	id, err := strconv.Atoi(mux.Vars(r)["id"])
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(id)
// 	price := prepareResponse()[id-1].Price
// 	points -= price/100 + 5

// 	jsonResponse, err := json.Marshal(PointsBalance{
// 		Points:  points,
// 		Balance: balance,
// 	})

// 	if err != nil {
// 		return
// 	}
// 	w.Write(jsonResponse)
// }

// func main() {
// 	router := mux.NewRouter()

// 	fmt.Println("run server")

// 	router.HandleFunc("/products", GetProduct).Methods("GET")
// 	router.HandleFunc("/state", GetState).Methods("GET")

// 	router.HandleFunc("/buy/{id:[0-9]+}", PostBuy).Methods("POST")
// 	router.HandleFunc("/redeem/{id:[0-9]+}", PostRedeem).Methods("POST")

// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedMethods:   []string{"GET", "POST"},
// 		AllowCredentials: true,
// 		Debug:            false,
// 	})

// 	handler := c.Handler(router)
// 	http.ListenAndServe(":4000", handler)
// }
