package service

import (
	"Golang-CRUD-API/models"
	"Golang-CRUD-API/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	return us.userRepository.GetAllUsers()
}

func (us *userService) GetUserByID(id uint) (*models.User, error) {
	return us.userRepository.GetUserByID(id)
}

func (us *userService) CreateUser(user *models.User) error {
	return us.userRepository.CreateUser(user)
}

func (us *userService) UpdateUser(user *models.User) error {
	return us.userRepository.UpdateUser(user)
}

func (us *userService) DeleteUser(id uint) error {
	user, err := us.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	return us.userRepository.DeleteUser(user)
}
