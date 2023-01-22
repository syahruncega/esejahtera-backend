package service

import (
	"kemiskinan/model"
	"kemiskinan/repository"
	"kemiskinan/request"
)

type UserService interface {
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	Create(userRequest request.CreateUserRequest) (model.User, error)
	Update(id int, userRequest request.UpdateUserRequest) (model.User, error)
	Delete(id int) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) FindAll() ([]model.User, error) {
	var users, err = s.userRepository.FindAll()

	return users, err
}

func (s *userService) FindById(id int) (model.User, error) {
	var user, err = s.userRepository.FindById(id)

	return user, err
}

func (s *userService) Create(userRequest request.CreateUserRequest) (model.User, error) {
	var user = model.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Email:    userRequest.Email,
		NoHp:     userRequest.NoHp,
		Role:     userRequest.Role,
	}

	var newUser, err = s.userRepository.Create(user)

	return newUser, err
}

func (s *userService) Update(id int, userRequest request.UpdateUserRequest) (model.User, error) {
	var user, err = s.userRepository.FindById(id)

	user.Username = userRequest.Username
	user.Password = userRequest.Password
	user.Email = userRequest.Email
	user.NoHp = userRequest.NoHp
	user.Role = userRequest.Role

	updatedUser, err := s.userRepository.Update(user)

	return updatedUser, err
}

func (s *userService) Delete(id int) (model.User, error) {
	var user, err = s.userRepository.FindById(id)

	deletedUser, err := s.userRepository.Delete(user)

	return deletedUser, err
}
