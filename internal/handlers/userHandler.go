package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/go-api/internal/models"
)

type UserUsecase interface {
	GetUsers() (*[]models.User, error)
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}

type UserHandler struct {
	Usercase UserUsecase
}

func NewUserHandler(uu UserUsecase) *UserHandler {
	return &UserHandler{Usercase: uu}
}

func (h *UserHandler) GetUsersHandler(c *fiber.Ctx) error {
	users, err := h.Usercase.GetUsers()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) GetUserByIdHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	user, err := h.Usercase.GetUserById(id)
	if err != nil {
		log.Fatal("User doesnt exist", err)
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal(err)
	}

	if err := h.Usercase.CreateUser(&user); err != nil {
		log.Fatal("Failed to creat:", err)
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		log.Fatal("Failed parsing body:", err)
	}

	if err := h.Usercase.UpdateUser(id, &user); err != nil {
		log.Fatal("Failed to update:", err)
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Fatal(err)
	}

	if err := h.Usercase.DeleteUser(id); err != nil {
		log.Fatal("Failed to delete:", err)
	}

	return c.SendStatus(200)
}
