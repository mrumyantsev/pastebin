package user

import "github.com/mrumyantsev/pastebin/internal/pkg/core"

type UserUseCase interface {
	CreateUser(input UserInput) (IdOutput, error)
	GetUsers(pg core.Pagination) ([]UserOutput, error)
	GetUser(id int) (UserOutput, error)
	UpdateUser(id int, usr UserInput) error
	DeleteUser(id int) error

	IsUserExists(username string) (IsExistsOutput, error)
	IsEmailExists(email string) (IsExistsOutput, error)
	UserCount() (CountOutput, error)
}
