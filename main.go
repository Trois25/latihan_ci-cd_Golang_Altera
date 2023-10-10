package main

import (
	"belajar-go-echo/app/configs"
	"belajar-go-echo/app/database"
	"belajar-go-echo/app/migration"
	"belajar-go-echo/app/router"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// db, err := configs.InitConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// err = migration.InitMigrationMysql(db)
	// if err != nil {
	// 	panic(err)
	// }

	// app := echo.New()
	// app.GET("/users", controller.GetAllUsers(db))
	// app.POST("/users", controller.CreateUser(db))
	// app.Start(":8080")

	cfg := configs.InitConfig()
	db := database.InitDBMysql(cfg)
	migration.InitMigrationMysql(db)

	e := echo.New()

	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
