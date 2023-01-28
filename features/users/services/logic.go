package services

import (
	"errors"
	"restapi-gofiber/features/users"
)

type usersService struct {
	userData users.RepositoryInterface
}

func NewUsersServices(data users.RepositoryInterface) users.ServicesInterface {
	return &usersService{
		data,
	}
}

func (services *usersService) GetAllUsers() (data []users.UserCore, err error) {
	data, err = services.userData.SelectAllUsers()
	return data, err
}

func (services *usersService) GetUserById(id int) (data users.UserCore, err error) {
	data, err = services.userData.SelectUserById(id)
	if err != nil {
		return users.UserCore{}, err
	} else if data.ID == 0 {
		return users.UserCore{}, errors.New("user not found")
	} else { // data.ID != 0
		return data, err
	}
}
func (services *usersService) PostUser(data users.UserCore) (row int, err error) {
	if data.Name == "" || data.Email == "" || data.Address == "" {
		return -1, errors.New("data cant be empty")
	}
	row, err = services.userData.CreateUser(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (services *usersService) PutUser(data users.UserCore) (row int, err error) {
	row, err = services.userData.UpdateUser(data)
	if err != nil {
		return -1, err
	}
	return row, err
}

func (services *usersService) DeleteUser(id int) (row int, err error) {
	row, err = services.userData.DeleteUser(id)
	if err != nil {
		return -1, err
	}
	return row, err
}
