package seller_controller

import (
	"arabian-snooker/models"

	"github.com/gin-gonic/gin"
)

type SellerController interface {
	CreateRates(c *gin.Context, req models.CreateMatch) (models.Matches, error)
	UpdateRate(c *gin.Context, req models.UpdateMatch) (models.Matches, error)
}
