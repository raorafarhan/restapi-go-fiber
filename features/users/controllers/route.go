package controllers

import (
	"github.com/gofiber/fiber/v2"
	"restapi-gofiber/features/users"
	"restapi-gofiber/utils/helper"
	"strconv"
)

type userControllers struct {
	userServices users.ServicesInterface
}

func New(f *fiber.App, services users.ServicesInterface) {
	controller := &userControllers{
		userServices: services,
	}
	f.Get("/users", controller.GetAllUsers)
	f.Get("/users/:id", controller.GetUserById)
	f.Post("/users", controller.PostUser)
	f.Put("/users/:id", controller.UpdateUser)
	f.Delete("/users/:id", controller.DeleteUser)
}

func (controller *userControllers) GetAllUsers(c *fiber.Ctx) error {

	data, err := controller.userServices.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return helper.GetResponseSuccess(c, data)
}
func (controller *userControllers) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	idConv, _ := strconv.Atoi(id)

	data, err := controller.userServices.GetUserById(idConv)
	if err != nil {
		return helper.GetResponseFailed(c)
	}
	return helper.GetResponseSuccess(c, data)
}
func (controller *userControllers) PostUser(c *fiber.Ctx) error {
	var data UserRequest
	errBind := c.BodyParser(&data)

	if errBind != nil {
		return helper.GetResponseFailedParseData(c)
	}
	row, err := controller.userServices.PostUser(ToCore(data))
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"message": err.Error(),
		})
	}
	if row != 1 {
		return c.Status(400).JSON(map[string]interface{}{"message": "failed to register"})
	}
	return helper.GetResponseSuccess(c, nil)
}

func (controller *userControllers) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idConv, _ := strconv.Atoi(id)

	if id == "" {
		return helper.GetResponseFailed(c)
	}
	var data UserRequest
	errBind := c.BodyParser(&data)

	if errBind != nil {
		return helper.GetResponseFailedParseData(c)
	}
	updateCore := ToCore(data)
	updateCore.ID = uint(idConv)

	row, err := controller.userServices.PutUser(updateCore)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": "data not found",
		})
	}

	if row != 1 {
		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": err,
		})
	}
	return helper.GetResponseSuccess(c, nil)
}

func (controller *userControllers) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idConv, _ := strconv.Atoi(id)
	if id == "" {
		return helper.GetResponseFailed(c)
	}
	row, err := controller.userServices.DeleteUser(idConv)
	if err != nil {
		return helper.GetResponseFailed(c)
	}
	if row != 1 {
		return helper.GetResponseFailed(c)
	}
	return helper.GetResponseSuccess(c, nil)
}
