package router

import (
	"arabian-snooker/controllers/admin_controller"
	"arabian-snooker/controllers/auth_service"
	"arabian-snooker/controllers/seller_controller"
	Validation "arabian-snooker/validation"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine           *gin.Engine
	AuthService      auth_service.AuthService
	Val              Validation.ValidationService
	Admin            admin_controller.AdminControllers
	SellerController seller_controller.SellerController
}

func NewRouter(

	authService auth_service.AuthService,
	valService Validation.ValidationService,
	AdminController admin_controller.AdminControllers,
	sellerController seller_controller.SellerController,
) *Router {

	engine := gin.Default()

	router := &Router{
		Engine:      engine,
		AuthService: authService,
		Val:         valService,
		Admin:       AdminController,

		SellerController: sellerController,
	}
	router.SetupRoutes()
	return router
}
