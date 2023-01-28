package controllers

import (
	"restapi-gofiber/features/users"
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromCore(data users.UserCore) UserResponse {
	return UserResponse{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Address:   data.Address,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromCoreList(data []users.UserCore) []UserResponse {
	var list []UserResponse
	for _, v := range data {
		list = append(list, FromCore(v))
	}
	return list
}
