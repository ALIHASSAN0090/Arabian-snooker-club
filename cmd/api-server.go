package cmd

import (
	"arabian-snooker/config"
	AdminControllerImpl "arabian-snooker/controllers/admin_controller/admin_impl"
	AuthServiceImpl "arabian-snooker/controllers/auth_service/auth_impl"
	SellerControllerImpl "arabian-snooker/controllers/seller_controller/seller_impl"
	database "arabian-snooker/database"
	dao "arabian-snooker/database/dao"
	"arabian-snooker/middleware"
	"arabian-snooker/router"
	Validation "arabian-snooker/validation"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var ApiServerCommand = &cobra.Command{
	Use:   "api",
	Short: "Api Starts Server",
	Run:   ExecuteApi,
}

func ExecuteApi(cmd *cobra.Command, args []string) {

	if err := config.LoadConfig(); err != nil {
		fmt.Println("Failed to load config: ", err)
	}

	fmt.Println(config.Cfg.Host, "PG_HOST")

	postgresDB, err := database.ConnectToPostgres()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to postgres!")
	}
	defer func() {
		fmt.Println("Closing database connection")
		postgresDB.Close()
	}()

	// migration := migrator.NewMigration()
	// if err := migration.RunMigrations(postgresDB); err != nil {
	// 	log.Fatalf("Migrations Failed %v", err)
	// }

	ValidationService := Validation.NewValidationService()

	Dao := dao.NewDao(postgresDB)

	middleware.Db = Dao

	fmt.Println("Starting Api Server")

	AuthService := AuthServiceImpl.NewAuthService(AuthServiceImpl.NewAuthServiceImpl{

		Dao: Dao,
		DB:  postgresDB,
	})
	AdminController := AdminControllerImpl.NewAdminController(AdminControllerImpl.NewAdminControllerImpl{

		Dao: Dao,
	})
	SellerController := SellerControllerImpl.NewSellerImpl(SellerControllerImpl.SellerController{
		Dao: Dao,
		DB:  postgresDB,
	})

	router := router.NewRouter(AuthService, ValidationService, AdminController, SellerController)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGABRT, os.Interrupt)

	Server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Cfg.Port),
		Handler: router.Engine,
	}

	go func() {
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {

		}
	}()

	<-stop

	fmt.Println("Stopping API server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := Server.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}

	<-ctx.Done()

	fmt.Println("API server stopped.")
}
