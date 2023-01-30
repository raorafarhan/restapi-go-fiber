package helper

import (
	"github.com/gofiber/fiber/v2"
)

type BaseResponseSuccess struct {
	Status  string      `json:"status"`
	Massage string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type BaseResponseFailed struct {
	Status  string `json:"status"`
	Massage string `json:"message"`
}

func GetResponseSuccess(c *fiber.Ctx, data interface{}) error {
	response := BaseResponseSuccess{}
	response.Status = "success"
	response.Massage = "success"
	response.Data = data

	return c.JSON(response)
}
func GetResponseFailed(c *fiber.Ctx) error {
	response := BaseResponseFailed{}
	response.Status = "failed"
	response.Massage = "data not found"

	return c.Status(404).JSON(response)
}

func GetResponseFailedParseData(c *fiber.Ctx) error {
	response := BaseResponseFailed{}
	response.Status = "failed"
	response.Massage = "data can't be empty"

	return c.Status(500).JSON(response)
}
