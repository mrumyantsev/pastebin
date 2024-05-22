package service

import (
	"github.com/mrumyantsev/pastebin/internal/domain/user"
	"github.com/mrumyantsev/pastebin/internal/pkg/core"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepo user.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(usr user.User) (int, error) {
	return s.userRepository.CreateUser(usr)
}

func (s *UserService) GetUsers(pg core.Pagination) ([]user.User, error) {
	return s.userRepository.GetUsers(pg)
}

func (s *UserService) GetUser(id int) (user.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *UserService) UpdateUser(id int, usr user.User) error {
	return s.userRepository.UpdateUser(id, usr)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}

func (s *UserService) IsUserExistsByUsername(username string) (bool, error) {
	return s.userRepository.IsUserExistsByUsername(username)
}

func (s *UserService) IsUserExistsById(id int) (bool, error) {
	return s.userRepository.IsUserExistsById(id)
}

func (s *UserService) IsEmailExists(email string) (bool, error) {
	return s.userRepository.IsEmailExists(email)
}

func (s *UserService) UserCount() (int, error) {
	return s.userRepository.UserCount()
}
