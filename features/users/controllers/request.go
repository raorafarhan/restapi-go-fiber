package controllers

import (
	"restapi-gofiber/features/users"
)

type UserRequest struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

func ToCore(data UserRequest) users.UserCore {
	return users.UserCore{
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}
}
