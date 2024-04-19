package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"haemong-be/controller"
	"haemong-be/repository"
	"haemong-be/service"
)

func main() {
	fmt.Println("hello hanul")
	app := fiber.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	user := app.Group("/user")
	user.Post("/sign-up", userController.UserController)

}
