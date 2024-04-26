package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"haemong-be/controller"
	"haemong-be/repository"
	"haemong-be/service"
	"os"
)

func main() {
	app := fiber.New()
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/haemong")
	if err != nil {
		fmt.Println("db boom..! ", err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	user := app.Group("/user")
	user.Post("/sign-up", userController.UserController)

	_ = app.Listen(":8080")
}
