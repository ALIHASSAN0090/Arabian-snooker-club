package seller_impl

import (
	"arabian-snooker/controllers/seller_controller"
	Dao "arabian-snooker/database/dao"
	"arabian-snooker/models"
	"arabian-snooker/utils"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type SellerControllerImpl struct {
	Dao Dao.Dao
	db  *sql.DB
}

type SellerController struct {
	Dao Dao.Dao
	DB  *sql.DB
}

func NewSellerImpl(input SellerController) seller_controller.SellerController {
	return &SellerControllerImpl{
		Dao: input.Dao,
		db:  input.DB,
	}
}

func (s *SellerControllerImpl) CreateRates(c *gin.Context, req models.CreateMatch) (models.Matches, error) {

	Id, err := utils.GetContextId(c)
	if err != nil || Id == 0 {
		return models.Matches{}, err
	}

	req.CreatedBy = Id

	match, err := s.Dao.CreateRates(req)
	if err != nil {
		return models.Matches{}, err
	}

	return match, nil

}

func (s *SellerControllerImpl) UpdateRate(c *gin.Context, req models.UpdateMatch) (models.Matches, error) {
	id, err := utils.GetContextId(c)
	utils.HandleError(err)

	req.UpdatedBy = id

	return s.Dao.UpdateRate(req)

}
