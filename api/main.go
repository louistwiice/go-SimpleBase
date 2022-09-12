package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louistwiice/go/simplebase/api/controllers"
	"github.com/louistwiice/go/simplebase/configs"
	"github.com/louistwiice/go/simplebase/repository"
	"github.com/louistwiice/go/simplebase/usecase/user"
)

// To load .env file
func init() {
	configs.Initialize()
}

func main() {
	log.Println("Server starting ...")

	// Start by connect to database
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", configs.GetString("MYSQL_USER"), configs.GetString("MYSQL_PASSWORD"), configs.GetString("MYSQL_HOST"), configs.GetString("MYSQL_DATABASE"))

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Panic("Database error: ... ", err.Error())
	}
	defer db.Close()

	// Connecting the database to the controllers and services
	userRepo := repository.NewUserSQL(db)
	userService := user.NewUSerService(userRepo)
	userController := controllers.NewUserController(userService)

	app := gin.Default()
	api_v1 := app.Group("api/v1")

	userController.MakeUserHandlers(api_v1.Group("user/"))

	app.Run(configs.GetString("SERVER_PORT"))
}
