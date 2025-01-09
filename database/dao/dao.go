package Dao

import (
	"arabian-snooker/models"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Dao interface {
	// Existing methods remain unchanged
	SignUp(req *models.Users) (*models.Users, error)
	CheckUserExists(req string) (bool, error)
	GetUser(req *models.LoginReq) (models.Users, error)
	IsActive(c *gin.Context, seller_id int64) (bool, error)
	CreateRates(req models.CreateMatch) (models.Matches, error)
	UpdateRate(req models.UpdateMatch) (models.Matches, error)

	// Corrected and Complete CRUD operations for all models

	// Users CRUD operations (assuming these are needed based on the context provided)
	// 	CreateUser(user *models.Users) (*models.Users, error)
	// 	GetUserByID(id int64) (*models.Users, error)
	// 	UpdateUser(user *models.Users) error
	// 	DeleteUser(id int64) error

	// 	// Matches CRUD operations
	// 	CreateMatch(match *models.Matches) (*models.Matches, error)
	// 	GetMatchByID(id int64) (*models.Matches, error)
	// 	UpdateMatch(match *models.Matches) error
	// 	DeleteMatch(id int64) error

	// 	// MatchesPlayed CRUD operations
	// 	CreateMatchesPlayed(mp *models.MatchesPlayed) (*models.MatchesPlayed, error)
	// 	GetMatchesPlayedByID(id int64) (*models.MatchesPlayed, error)
	// 	UpdateMatchesPlayed(mp *models.MatchesPlayed) error
	// 	DeleteMatchesPlayed(id int64) error

	// 	// Centuries CRUD operations
	// 	CreateCenturies(c *models.Centuries) (*models.Centuries, error)
	// 	GetCenturiesByID(id int64) (*models.Centuries, error)
	// 	UpdateCenturies(c *models.Centuries) error
	// 	DeleteCenturies(id int64) error

	// // CenturiesPlayed CRUD operations
	// CreateCenturiesPlayed(cp *models.CenturiesPlayed) (*models.CenturiesPlayed, error)
	// GetCenturiesPlayedByID(id int64) (*models.CenturiesPlayed, error)
	// UpdateCenturiesPlayed(cp *models.CenturiesPlayed) error
	// DeleteCenturiesPlayed(id int64) error
}

func NewDao(db *sql.DB) Dao {
	return &DaoImp{
		db: db,
	}
}

type DaoImp struct {
	db *sql.DB
}
