package users

import "time"

type UserCore struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServicesInterface interface {
	GetAllUsers() (data []UserCore, err error)
	GetUserById(id int) (data UserCore, err error)
	PostUser(data UserCore) (row int, err error)
	PutUser(data UserCore) (row int, err error)
	DeleteUser(id int) (row int, err error)
}
type RepositoryInterface interface {
	SelectAllUsers() (data []UserCore, err error)
	SelectUserById(id int) (data UserCore, err error)
	CreateUser(data UserCore) (row int, err error)
	UpdateUser(data UserCore) (row int, err error)
	DeleteUser(id int) (row int, err error)
}
