package router

import (
	"arabian-snooker/models"
	"arabian-snooker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) CreateRates(c *gin.Context) {

	var matches models.CreateMatch

	if err := c.ShouldBindJSON(&matches); err != nil {
		utils.HandleError(err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				Message:    "invalid request data",
				Detail:     err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		})
		return
	}
	if err := r.Val.ValidateMatches(&matches); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				Message:    "validation error",
				Detail:     err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		})
		return
	}

	data, err := r.SellerController.CreateRates(c, matches)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				Message:    "error in creating rates",
				Detail:     err.Error(),
				StatusCode: http.StatusInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
