package validation_impl

import (
	"arabian-snooker/models"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ValidationServiceImpl struct{}

func (vs *ValidationServiceImpl) ValidateEmailPassword(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value == "" {
		return false
	}
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", value)
	return match
}

func (vs *ValidationServiceImpl) ValidateReq(c *gin.Context, request interface{}) []string {
	validate := validator.New()

	validate.RegisterValidation("alphanum", func(fl validator.FieldLevel) bool {
		return vs.ValidateEmailPassword(fl)
	})

	var errorMsgs []string
	if err := validate.Struct(request); err != nil {
		ValidationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range ValidationErrors {
			fieldName := fieldError.Field()
			tag := fieldError.Tag()
			errorMsg := fmt.Sprintf("%s is invalid. Error: %s", fieldName, tag)
			errorMsgs = append(errorMsgs, errorMsg)
		}
		return errorMsgs
	}
	return errorMsgs
}

func (vs *ValidationServiceImpl) ValidateMatches(match *models.CreateMatch) error {

	if match.MatchName != strings.ToLower(match.MatchName) {
		return errors.New("match name must be in lowercase")
	}

	if match.PerMatchPrice < 0 {
		return errors.New("per match price cannot be negative")
	}

	return nil
}

func (vs *ValidationServiceImpl) ValidateUpdateMatch(m *models.UpdateMatch) error {
	if m.ID <= 0 {
		return errors.New("ID must be greater than 0")
	}
	return nil
}
