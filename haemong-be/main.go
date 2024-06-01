package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"haemong-be/controller"
	"haemong-be/repository"
	"haemong-be/service"
	"os"
)

func main() {
	app := fiber.New()
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	gptkey := os.Getenv("GPT")
	db, err := sql.Open("mysql", "admin:"+password+"@tcp("+host+":3306)/haemong")
	if err != nil {
		fmt.Println("db boom..! ", err)
		return
	}

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	userRepository := repository.NewUserRepository(db)
	feedRepository := repository.NewFeedRepository(db)
	authRepository := repository.NewAuthRepository(db)
	ilgiRepository := repository.NewIlgiRepository(db)
	commentRepository := repository.NewCommentRepository(db)

	userService := service.NewUserService(userRepository)
	feedService := service.NewFeedService(feedRepository, userRepository)
	authService := service.NewAuthService(authRepository, userRepository)
	ilgiService := service.NewIlgiService(ilgiRepository)
	commentService := service.NewCommentService(commentRepository)
	aiService := service.NewAiService(gptkey)

	userController := controller.NewUserController(userService)
	feedController := controller.NewFeedController(feedService)
	authController := controller.NewAuthController(authService)
	ilgiController := controller.NewIlgiController(ilgiService)
	commentController := controller.NewCommentController(commentService)
	aiController := controller.NewAiController(aiService)

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
	feed.Get("/", feedController.ReadAllFeeds)
	feed.Post("/like/:feedId", feedController.AddLike)
	feed.Post("/dislike/:feedId", feedController.RemoveLike)

	comment := app.Group("/comment")
	comment.Get("/:feedId", commentController.ReadCommentsOnTheFeed)
	comment.Post("/", commentController.CreateComment)

	ai := app.Group("/ai")
	ai.Get("/", aiController.Haemong)

	_ = app.Listen(":8080")
}
