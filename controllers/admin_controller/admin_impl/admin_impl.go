package admin_controller_impl

import (
	"arabian-snooker/controllers/admin_controller"
	Dao "arabian-snooker/database/dao"
)

type AdminControllerImpl struct {
	Dao Dao.Dao
}

func NewAdminController(input NewAdminControllerImpl) admin_controller.AdminControllers {
	return &AdminControllerImpl{

		Dao: input.Dao,
	}
}

type NewAdminControllerImpl struct {
	Dao Dao.Dao
}
