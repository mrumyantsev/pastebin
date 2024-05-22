package user

import "github.com/mrumyantsev/pastebin/internal/pkg/core"

type UserRepository interface {
	CreateUser(usr User) (int, error)
	GetUsers(pg core.Pagination) ([]User, error)
	GetUser(id int) (User, error)
	UpdateUser(id int, usr User) error
	DeleteUser(id int) error

	IsUserExistsByUsername(username string) (bool, error)
	IsUserExistsById(id int) (bool, error)
	IsEmailExists(email string) (bool, error)
	UserCount() (int, error)
}
