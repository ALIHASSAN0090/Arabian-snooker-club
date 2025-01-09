package router

import (
	"arabian-snooker/models"
	"arabian-snooker/utils"
	"fmt"
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, models.SuccessResponse{
		GinContext: c,
		Data:       data,
		StatusCode: http.StatusOK,
		Message:    "Rates created successfully",
	})
}

func (r *Router) UpdateRate(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Print(idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				Message:    "invalid match id",
				Detail:     "The provided id '" + idStr + "' is not a valid integer",
				StatusCode: http.StatusBadRequest,
			},
		})
		return
	}

	var update models.UpdateMatch
	update.ID = id

	if err := c.ShouldBindJSON(&update); err != nil {
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

	if err := r.Val.ValidateUpdateMatch(&update); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				Message:    "validation error",
				Detail:     err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		})
		return
	}

	updatedData, err := r.SellerController.UpdateRate(c, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				Message:    "error in updating rates",
				Detail:     err.Error(),
				StatusCode: http.StatusInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		GinContext: c,
		Data:       updatedData,
		StatusCode: http.StatusOK,
		Message:    "Rates updated successfully",
	})
}
