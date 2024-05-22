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
	feedRepository := repository.NewFeedRepository(db)
	authRepository := repository.NewAuthRepository(db)
	ilgiRepository := repository.NewIlgiRepository(db)

	userService := service.NewUserService(userRepository)
	feedService := service.NewFeedService(feedRepository, userRepository)
	authService := service.NewAuthService(authRepository, userRepository)
	ilgiService := service.NewIlgiService(ilgiRepository)

	userController := controller.NewUserController(userService)
	feedController := controller.NewFeedController(feedService)
	authController := controller.NewAuthController(authService)
	ilgiController := controller.NewIlgiController(ilgiService)

	user := app.Group("/user")
	user.Post("/sign-up", userController.SignUp)
	user.Post("/change-password", userController.ChangePassword)
	user.Delete("/withdrawal", userController.Withdrawal)

	auth := app.Group("/auth")
	auth.Post("/sign-in", authController.SignIn)

	ilgi := app.Group("/ilgi")
	ilgi.Post("/save", ilgiController.SaveIlgi)
	ilgi.Put("/modify", ilgiController.ModifyIlgi)
	ilgi.Delete("/:id", ilgiController.DeleteIlgi)
	ilgi.Get("/:search", ilgiController.SearchIlgi)

	feed := app.Group("/feed")
	feed.Post("/", feedController.SaveFeed)

	_ = app.Listen(":8080")
}
