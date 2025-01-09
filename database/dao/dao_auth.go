package Dao

import (
	"arabian-snooker/models"

	"github.com/gin-gonic/gin"
)

func (a *DaoImp) CheckUserExists(req string) (bool, error) {

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := a.db.QueryRow(checkQuery, req).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (a *DaoImp) SignUp(req *models.Users) (*models.Users, error) {

	query := `INSERT INTO users(name, email, password, created_at) 
	VALUES($1, $2, $3, NOW()) RETURNING id`
	err := a.db.QueryRow(query, req.Name, req.Email, req.Password).Scan(&req.ID)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (a *DaoImp) GetUser(req *models.LoginReq) (models.Users, error) {
	var user models.Users

	query := `
	SELECT id,name, email,password,role
    FROM users WHERE email = $1
	`
	err := a.db.QueryRow(query, req.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func (s *DaoImp) IsActive(c *gin.Context, seller_id int64) (bool, error) {
	query := `
	SELECT active
	FROM users
	WHERE id = $1
	`

	var active bool
	err := s.db.QueryRow(query, seller_id).Scan(&active)
	if err != nil {
		return false, err
	}

	return active, nil
}
