package auth_service

import (
	"arabian-snooker/models"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, req *models.Users) (*models.Users, string, error)
	CheckUserExists(req *models.Users) (bool, error)
	ProcessLogin(ctx *gin.Context, req *models.LoginReq) (string, error)
}
