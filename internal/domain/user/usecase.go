package user

import "github.com/mrumyantsev/pastebin/internal/pkg/core"

type UserUseCase interface {
	CreateUser(input UserInput) error
	GetUsers(pg core.Pagination) ([]UserOutput, error)
	GetUser(id int) (UserOutput, error)
	UpdateUser(id int, usr UserInput) error
	DeleteUser(id int) error
}
