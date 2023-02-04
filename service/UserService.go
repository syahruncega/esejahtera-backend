package service

import (
	"kemiskinan/helper"
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
	Login(username string) (model.User, error)
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
	var password = userRequest.Password
	var hashedPassword, _ = helper.HashPassword(password)

	var user = model.User{
		Username: userRequest.Username,
		Password: hashedPassword,
		Email:    userRequest.Email,
		NoHp:     userRequest.NoHp,
		Role:     userRequest.Role,
	}

	var newUser, err = s.userRepository.Create(user)

	return newUser, err
}

func (s *userService) Update(id int, userRequest request.UpdateUserRequest) (model.User, error) {
	var user, err = s.userRepository.FindById(id)

	var password = userRequest.Password
	var hashedPassword, _ = helper.HashPassword(password)

	user.Username = userRequest.Username
	user.Password = hashedPassword
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

func (s *userService) Login(username string) (model.User, error) {
	var user, err = s.userRepository.Login(username)

	return user, err
}
