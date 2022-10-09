package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func prepareResponse() []Product {
	var products []Product
	var p Product

	p.Id = 1
	p.Name = "ball"
	p.Price = 10_000
	p.Image = "https://laika.com.co/cdn-cgi/image/onerror=redirect,format=auto,fit=scale-down,width=600,quality=80/https://laikapp.s3.amazonaws.com/dev_images_products/11145_169454_Argos___Juguete_Surtidos_Mediano_1657722486_500x500.jpg"
	products = append(products, p)

	p.Id = 2
	p.Name = "snaks pot"
	p.Price = 20_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/images_products/1_4328_SUPER_ORANGE_402X385.jpg"
	products = append(products, p)

	p.Id = 3
	p.Name = "rope"
	p.Price = 30_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/74104_Beco_Pets___Aro_De_C____amo_1638208176_0_300x300.png"
	products = append(products, p)

	p.Id = 4
	p.Name = "pillow"
	p.Price = 40_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/23370_136193_Indupet_Peluche_Galleta_1644869825_500x500.jpg"
	products = append(products, p)

	p.Id = 5
	p.Name = "buffalo"
	p.Price = 50_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/15675_168422_Funimals___B__falo_Con_Sonido__1657228119_500x500.png"
	products = append(products, p)

	p.Id = 6
	p.Name = "pineapple"
	p.Price = 60_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/23378_136460_Indupet___Peluche_Pi__a__1644876493_500x500.jpg"
	products = append(products, p)

	p.Id = 7
	p.Name = "koala"
	p.Price = 70_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/67207_kong_peluche_snuzzies_koala__1622842429_0_1000x1000.jpg"
	products = append(products, p)

	p.Id = 8
	p.Name = "steak"
	p.Price = 80_000
	p.Image = "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/23041_136031_Indupet___Peluche_Hueso_Osobuco__1644866425_500x500.jpg"
	products = append(products, p)

	return products
}

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

type PointsBalance struct {
	Points  int `json:"points"`
	Balance int `json:"balance"`
}

var points = 0
var balance = 200_000

func GetProduct(w http.ResponseWriter, r *http.Request) {
	products := prepareResponse()
	jsonResponse, err := json.Marshal(products)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func GetState(w http.ResponseWriter, r *http.Request) {
	jsonResponse, err := json.Marshal(PointsBalance{
		Points:  points,
		Balance: balance,
	})

	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func PostBuy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	price := prepareResponse()[id-1].Price

	points += price/1000 + 5
	balance = balance - price

	jsonResponse, err := json.Marshal(PointsBalance{
		Points:  points,
		Balance: balance,
	})

	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func PostRedeem(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	price := prepareResponse()[id-1].Price
	points -= price/100 + 5

	jsonResponse, err := json.Marshal(PointsBalance{
		Points:  points,
		Balance: balance,
	})

	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func main() {
	router := mux.NewRouter()

	fmt.Println("run server")

	router.HandleFunc("/products", GetProduct).Methods("GET")
	router.HandleFunc("/state", GetState).Methods("GET")

	router.HandleFunc("/buy/{id:[0-9]+}", PostBuy).Methods("POST")
	router.HandleFunc("/redeem/{id:[0-9]+}", PostRedeem).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowCredentials: true,
		Debug:            false,
	})

	handler := c.Handler(router)
	http.ListenAndServe(":4000", handler)
}
