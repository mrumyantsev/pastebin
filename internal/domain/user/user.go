package user

import (
	"time"

	"github.com/mrumyantsev/pastebin/internal/pkg/core"
)

const (
	TabUsers = "users"

	AllCols = "*"

	ColId           = "id"
	ColUsername     = "username"
	ColPasswordHash = "password_hash"
	ColFirstName    = "first_name"
	ColLastName     = "last_name"
	ColEmail        = "email"
	ColCreatedAt    = "created_at"
)

type UserInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserOutput struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	CreatedAt    string `json:"createdAt"`
}

func NewUserOutput(
	id int,
	username string,
	passwordHash string,
	firstName string,
	lastName string,
	email string,
	createdAt time.Time,
) UserOutput {
	return UserOutput{
		Id:           id,
		Username:     username,
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		CreatedAt:    createdAt.Format(core.TimeFormat),
	}
}

type IdOutput struct {
	Id int `json:"id"`
}

func NewIdOutput(id int) IdOutput {
	return IdOutput{Id: id}
}

type CountOutput struct {
	Count int `json:"count"`
}

func NewCountOutput(count int) CountOutput {
	return CountOutput{Count: count}
}

type IsExistsOutput struct {
	IsExists bool `json:"isExists"`
}

func NewIsExistsOutput(cond bool) IsExistsOutput {
	return IsExistsOutput{IsExists: cond}
}

type User struct {
	Id           int       `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Email        string    `db:"email"`
	CreatedAt    time.Time `db:"created_at"`
}

func NewUser(
	id int,
	username string,
	passwordHash string,
	firstName string,
	lastName string,
	email string,
) User {
	return User{
		Id:           id,
		Username:     username,
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		CreatedAt:    time.Now(),
	}
}
