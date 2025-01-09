package seller_impl

import (
	"arabian-snooker/controllers/seller_controller"
	Dao "arabian-snooker/database/dao"
	"database/sql"
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
