package usecase

import (
	"github.com/mrumyantsev/pastebin/internal/domain/user"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/core"
	"github.com/mrumyantsev/pastebin/internal/pkg/passhash"
	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
)

type UserUseCase struct {
	config      *config.Config
	userService user.UserService
}

func NewUserUseCase(cfg *config.Config, userService user.UserService) *UserUseCase {
	return &UserUseCase{config: cfg, userService: userService}
}

func (u *UserUseCase) CreateUser(input user.UserInput) error {
	if core.IsMissingInput(input.Username, input.Password, input.FirstName, input.LastName, input.Email) {
		return errlib.NewCustomError(core.MsgInputMissing)
	}

	if core.IsInputLengthTooLong(user.MaxUsernameLength, input.Username) {
		return errlib.NewCustomError(core.MsgInputExceeds)
	}

	if core.IsInputLengthTooLong(user.MaxInputLength, input.Password, input.FirstName, input.LastName, input.Email) {
		return errlib.NewCustomError(core.MsgInputExceeds)
	}

	exists, err := u.userService.IsUserExists(input.Username)
	if err != nil {
		return err
	}
	if exists {
		return errlib.NewCustomError(core.MsgExists)
	}

	if exists, err = u.userService.IsEmailExists(input.Email); err != nil {
		return err
	}
	if exists {
		return errlib.NewCustomError(core.MsgEmailExists)
	}

	usr := user.NewUser(
		0,
		input.Username,
		passhash.PasswordHash(input.Password, u.config.PasswordHashSalt),
		input.FirstName,
		input.LastName,
		input.Email,
	)

	return u.userService.CreateUser(usr)
}

func (u *UserUseCase) GetUsers(pg core.Pagination) ([]user.UserOutput, error) {
	users, err := u.userService.GetUsers(pg)
	if err != nil {
		return nil, err
	}

	usersLength := len(users)

	usersOutput := make([]user.UserOutput, usersLength)

	for i := 0; i < usersLength; i++ {
		usersOutput[i] = user.NewUserOutput(
			users[i].Id,
			users[i].Username,
			users[i].PasswordHash,
			users[i].FirstName,
			users[i].LastName,
			users[i].Email,
			users[i].CreatedAt,
		)
	}

	return usersOutput, nil
}

func (u *UserUseCase) GetUser(id int) (user.UserOutput, error) {
	var output user.UserOutput

	usr, err := u.userService.GetUser(id)
	if err != nil {
		return output, err
	}

	output = user.NewUserOutput(
		usr.Id,
		usr.Username,
		usr.PasswordHash,
		usr.FirstName,
		usr.LastName,
		usr.Email,
		usr.CreatedAt,
	)

	return output, nil
}

func (u *UserUseCase) UpdateUser(id int, input user.UserInput) error {
	usr := user.NewUser(
		id,
		input.Username,
		passhash.PasswordHash(input.Password, u.config.PasswordHashSalt),
		input.FirstName,
		input.LastName,
		input.Email,
	)

	if err := u.userService.UpdateUser(id, usr); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) DeleteUser(id int) error {
	if err := u.userService.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
