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

func (u *UserUseCase) CreateUser(input user.UserInput) (user.IdOutput, error) {
	var output user.IdOutput

	if core.IsInputMissing(input.Username, input.Password, input.FirstName, input.LastName, input.Email) {
		return output, errlib.NewCustomError(core.MsgInputMissing)
	}

	if core.IsInputExceeds(user.MaxUsernameLength, input.Username) {
		return output, errlib.NewCustomError(core.MsgInputExceeds)
	}

	if core.IsInputExceeds(user.MaxInputLength, input.Password, input.FirstName, input.LastName, input.Email) {
		return output, errlib.NewCustomError(core.MsgInputExceeds)
	}

	if !core.IsEmailValid(input.Email) {
		return output, errlib.NewCustomError(core.MsgEmailInvalid)
	}

	exists, err := u.userService.IsUserExistsByUsername(input.Username)
	if err != nil {
		return output, err
	}
	if exists {
		return output, errlib.NewCustomError(core.MsgExists)
	}

	if exists, err = u.userService.IsEmailExists(input.Email); err != nil {
		return output, err
	}
	if exists {
		return output, errlib.NewCustomError(core.MsgEmailExists)
	}

	usr := user.NewUser(
		0,
		input.Username,
		passhash.PasswordHash(input.Password, u.config.PasswordHashSalt),
		input.FirstName,
		input.LastName,
		input.Email,
	)

	id, err := u.userService.CreateUser(usr)
	if err != nil {
		return output, err
	}

	output = user.NewIdOutput(id)

	return output, nil
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

	exists, err := u.userService.IsUserExistsById(id)
	if err != nil {
		return output, err
	}
	if !exists {
		return output, errlib.NewCustomError(core.MsgNotExists)
	}

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
	exists, err := u.userService.IsUserExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errlib.NewCustomError(core.MsgNotExists)
	}

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
	exists, err := u.userService.IsUserExistsById(id)
	if err != nil {
		return err
	}
	if !exists {
		return errlib.NewCustomError(core.MsgNotExists)
	}

	if err := u.userService.DeleteUser(id); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) IsUserExists(username string) (user.IsExistsOutput, error) {
	var output user.IsExistsOutput

	exists, err := u.userService.IsUserExistsByUsername(username)
	if err != nil {
		return output, err
	}

	output = user.NewIsExistsOutput(exists)

	return output, nil
}

func (u *UserUseCase) IsEmailExists(email string) (user.IsExistsOutput, error) {
	var output user.IsExistsOutput

	exists, err := u.userService.IsEmailExists(email)
	if err != nil {
		return output, err
	}

	output = user.NewIsExistsOutput(exists)

	return output, nil
}

func (u *UserUseCase) UserCount() (user.CountOutput, error) {
	var output user.CountOutput

	count, err := u.userService.UserCount()
	if err != nil {
		return output, err
	}

	output = user.NewCountOutput(count)

	return output, nil
}
