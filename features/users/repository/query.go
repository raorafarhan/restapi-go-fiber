package repository

import (
	"restapi-gofiber/features/users"

	"gorm.io/gorm"
)

type dataUser struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.RepositoryInterface {
	return &dataUser{
		db,
	}
}

func (repo *dataUser) SelectAllUsers() ([]users.UserCore, error) {
	var users []User
	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	userCore := toCoreList(users)
	return userCore, nil
}

func (repo *dataUser) SelectUserById(id int) (users.UserCore, error) {
	var userList User
	userList.ID = uint(id)
	tx := repo.db.Where("id = ?", id).First(&userList)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	userData := userList.toCore()
	return userData, nil
}

func (repo *dataUser) CreateUser(data users.UserCore) (row int, err error) {
	var user User

	user.Name = data.Name
	user.Email = data.Email
	user.Address = data.Address

	tx := repo.db.Create(&user)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataUser) UpdateUser(data users.UserCore) (int, error) {
	var userUpdate User
	txDataOld := repo.db.First(&userUpdate, data.ID)

	if txDataOld.Error != nil {
		return -1, txDataOld.Error
	}

	if data.Name != "" {
		userUpdate.Name = data.Name
	}

	if data.Address != "" {
		userUpdate.Address = data.Address
	}

	tx := repo.db.Save(&userUpdate)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *dataUser) DeleteUser(id int) (row int, err error) {
	var user User
	user.ID = uint(id)
	tx := repo.db.Delete(&user)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}
