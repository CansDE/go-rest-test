package repository

import (
	"test-rest-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(user *models.User) {
	u.DB.Create(user)
}

func (u *UserRepository) GetAll() []models.User {
	var users []models.User
	u.DB.Find(&users)
	return users
}
