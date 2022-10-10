package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	points_repo "loyalty/internal/repositories/points"
	product_repo "loyalty/internal/repositories/product"
	user_repo "loyalty/internal/repositories/user"

	points_srv "loyalty/internal/core/services/points"
	product_srv "loyalty/internal/core/services/product"
	user_srv "loyalty/internal/core/services/user"

	points_hdl "loyalty/internal/handlers/points"
	product_hdl "loyalty/internal/handlers/product"
	user_hdl "loyalty/internal/handlers/user"

	points_evt "loyalty/internal/core/events/points"
)

func main() {

	pointsRepository := points_repo.New()
	productRepository := product_repo.New()
	userRepository := user_repo.New()

	pointsEvents := points_evt.New()

	pointsService := points_srv.New(pointsEvents, pointsRepository, productRepository)
	productService := product_srv.New(productRepository)
	userService := user_srv.New(userRepository)

	pointsHandler := points_hdl.New(pointsService)
	productHandler := product_hdl.New(productService)
	userHandler := user_hdl.New(userService)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/points/:id", pointsHandler.Get)
	router.GET("/products", productHandler.All)
	router.GET("/users/:id", userHandler.Get)
	router.POST("/buy/:id/:idProduct", pointsHandler.Buy)
	router.POST("/redeem/:id/:idProduct", pointsHandler.Redeem)

	router.Run(":8080")
}
