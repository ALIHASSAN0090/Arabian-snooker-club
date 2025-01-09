package auth_service_impl

import (
	"arabian-snooker/controllers/auth_service"
	dao "arabian-snooker/database/dao"
	"arabian-snooker/middleware"
	"arabian-snooker/models"
	"arabian-snooker/utils"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
	Dao dao.Dao
	db  *sql.DB
}

type NewAuthServiceImpl struct {
	Dao dao.Dao
	DB  *sql.DB
}

func NewAuthService(input NewAuthServiceImpl) auth_service.AuthService {
	return &AuthServiceImpl{

		Dao: input.Dao,
		db:  input.DB,
	}
}

func (a *AuthServiceImpl) SignUp(ctx *gin.Context, req *models.Users) (*models.Users, string, error) {

	exists, err := a.Dao.CheckUserExists(*req.Email)
	if err != nil {

		return &models.Users{}, "error in checking if user exist", err
	}

	if exists {
		return &models.Users{}, "User already exists", nil
	}
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {

		return &models.Users{}, "Error hashing password", err
	}

	passwordMatch, _ := utils.VerifyPassword(hashed, req.Password)
	if !passwordMatch {

		return &models.Users{}, "password is invalid", fmt.Errorf("password is invalid")
	}

	req.Password = hashed

	userCreated, err := a.Dao.SignUp(req)
	if err != nil {

		return &models.Users{}, "Error during SignUp", err
	}

	return userCreated, "Signed Up successfully", nil
}

func (a *AuthServiceImpl) CheckUserExists(req *models.Users) (bool, error) {

	return a.Dao.CheckUserExists(*req.Email)
}

func (a *AuthServiceImpl) ProcessLogin(ctx *gin.Context, req *models.LoginReq) (string, error) {

	fmt.Println("ProcessLogin called with email:", req.Email)
	exists, err := a.Dao.CheckUserExists(req.Email)
	if err != nil {
		fmt.Println("Error checking if user exists:", err)
		return "", err
	}

	fmt.Println("User exists:", exists)
	if !exists {
		fmt.Println("User does not exist")
		return "", fmt.Errorf("user does not exist")
	}

	user, err := a.Dao.GetUser(req)
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return "", err
	}

	fmt.Println("User retrieved:", user.Email)
	passwordMatch, _ := utils.VerifyPassword(user.Password, req.Password)
	fmt.Println("Password match result:", passwordMatch)
	if !passwordMatch {
		fmt.Println("Invalid credentials")
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := middleware.GenerateAccessToken(&user)
	if err != nil {
		fmt.Println("Error generating access token:", err)
		utils.HandleError(err)
		return "", err
	}

	fmt.Println("Access token generated successfully")
	return token, nil
}
