package base

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/in-mai-space/disco/entities/models"
	"github.com/in-mai-space/disco/types"
	"github.com/in-mai-space/disco/utils"
)

type UserController struct {
	types.BaseController[models.User]
	userService UserServiceInterface
}

func NewUserController(userService UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var requestBody models.UserCreateRequestBody
	if err := u.ParseDTO(c, &requestBody, "user"); err != nil {
		return u.HandleError(c, err)
	}

	user := &models.User{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	if err := utils.HashPassword(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	createdUser, err := u.userService.CreateUser(user)
	if err != nil {
		return u.HandleError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(createdUser)
}

func (u *UserController) GetUserByID(c *fiber.Ctx) error {
	id, err := u.ValidateID(c, "user")
	if err != nil {
		return u.HandleError(c, err)
	}

	user, err := u.userService.GetUserByID(id)
	if err != nil {
		return u.HandleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := u.ValidateID(c, "user")
	if err != nil {
		return u.HandleError(c, err)
	}

	var requestBody models.UserNameUpdateRequestBody
	if err := u.ParseDTO(c, &requestBody, "user"); err != nil {
		return u.HandleError(c, err)
	}

	user := &models.User{
		Name:     requestBody.Name,
	}

	_, err = u.userService.GetUserByID(id)
	if err != nil {
		return u.HandleError(c, err)
	}

	updatedUser, err := u.userService.UpdateUser(user, id)
	if err != nil {
		return u.HandleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := u.ValidateID(c, "user")
	if err != nil {
		return u.HandleError(c, err)
	}

	if err := u.userService.DeleteUser(id); err != nil {
		return u.HandleError(c, err)
	}

	return c.Status(http.StatusNoContent).Send(nil)
}
