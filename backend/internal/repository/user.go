package repository

import (
	"echo-api/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetByEmail(user *domain.User, email string) error {
	return ur.db.Where("email = ?", email).First(user).Error
}

func (ur *UserRepository) GetByID(user *domain.User, id uint) error {
	return ur.db.First(user, id).Error
}

func (ur *UserRepository) Create(user *domain.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) Update(user *domain.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(user *domain.User) error {
	return ur.db.Delete(user).Error
}
