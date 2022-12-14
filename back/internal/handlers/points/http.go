package points_hdl

import (
	"loyalty/internal/core/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	pointsService ports.PointsService
}

func New(pointsService ports.PointsService) *HTTPHandler {
	return &HTTPHandler{
		pointsService: pointsService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	point, err := hdl.pointsService.Get(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, point)
}

func (hdl *HTTPHandler) Redeem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idProduct, _ := strconv.Atoi(c.Param("idProduct"))
	points, _ := hdl.pointsService.Redeem(id, idProduct)
	c.JSON(200, points)
}

func (hdl *HTTPHandler) Buy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	idProduct, _ := strconv.Atoi(c.Param("idProduct"))
	points, _ := hdl.pointsService.Buy(id, idProduct)
	c.JSON(200, points)
}
