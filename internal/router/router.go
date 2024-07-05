package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/go-api/internal/handlers"
	"github.com/just-umyt/go-api/internal/mysqlapi"
	"github.com/just-umyt/go-api/internal/repository"
	"github.com/just-umyt/go-api/internal/usecase"
)

func NewApp() *fiber.App {
	app := fiber.New()

	dsn := "root:@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db := mysqlapi.InitDB(dsn)

	userRepo := repository.NewUserRepo(db)
	userUseCases := usecase.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCases)

	app.Get("/users", userHandler.GetUsersHandler)
	app.Get("/users/:id", userHandler.GetUserByIdHandler)
	app.Post("/users", userHandler.CreateUserHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/users/:id", userHandler.DeleteUserHandler)

	return app
}
