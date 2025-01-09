package Validation

import (
	"arabian-snooker/models"
	"arabian-snooker/validation/validation_impl"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ValidationService interface {
	ValidateReq(c *gin.Context, request interface{}) []string
	ValidateEmailPassword(fl validator.FieldLevel) bool
	ValidateMatches(match *models.CreateMatch) error
}

func NewValidationService() ValidationService {
	return &validation_impl.ValidationServiceImpl{}
}
