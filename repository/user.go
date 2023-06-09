package repository

import (
	"Golang-CRUD-API/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.db.Find(&users).Error
	return users, err
}

func (ur *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) UpdateUser(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *userRepository) DeleteUser(user *models.User) error {
	return ur.db.Delete(user).Error
}
