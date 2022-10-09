package product_hdl

import (
	"loyalty/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	productService ports.ProductService
}

func New(productService ports.ProductService) *HTTPHandler {
	return &HTTPHandler{
		productService: productService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := hdl.productService.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, product)
}
