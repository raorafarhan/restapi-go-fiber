package repository

import (
	"restapi-gofiber/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Email   string
	Address string
}

func fromCore(data users.UserCore) User {
	return User{
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}
}

func (data *User) toCore() users.UserCore {
	return users.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Address:   data.Address,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.UserCore {
	var list []users.UserCore
	for _, v := range data {
		list = append(list, v.toCore())
	}
	return list
}
