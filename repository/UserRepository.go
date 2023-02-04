package repository

import (
	"kemiskinan/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) (model.User, error)
	Login(username string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User

	var err = r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) FindById(id int) (model.User, error) {
	var user model.User

	var err = r.db.Take(&user, id).Error

	return user, err
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	var err = r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) Update(user model.User) (model.User, error) {

	var err = r.db.Model(&user).Updates(model.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		NoHp:     user.NoHp,
		Role:     user.Role,
	}).Error

	return user, err
}

func (r *userRepository) Delete(user model.User) (model.User, error) {
	var err = r.db.Delete(&user).Error

	return user, err
}

func (r *userRepository) Login(username string) (model.User, error) {
	var user model.User

	var err = r.db.Where("username = ?", username).Take(&user).Error

	return user, err
}
