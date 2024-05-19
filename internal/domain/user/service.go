package user

import "github.com/mrumyantsev/pastebin/internal/pkg/core"

type UserService interface {
	CreateUser(usr User) error
	GetUsers(pg core.Pagination) ([]User, error)
	GetUser(id int) (User, error)
	UpdateUser(id int, usr User) error
	DeleteUser(id int) error

	IsUserExists(username string) (bool, error)
	IsEmailExists(email string) (bool, error)
}
