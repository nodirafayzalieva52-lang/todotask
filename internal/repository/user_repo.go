package repository

import (
	"fmt"
	"nd/internal/models"

	"gorm.io/gorm"
)

type Iuserrepo interface {
  CreateUser(user models.Users) error
  GetAllUsers() ([]models.Users, error)
  GetByUserID(id uint) (models.Users, error)
  DeleteUser(id uint) error
  UpdateUsers(id uint) error
}

type usersrepo struct {
  db *gorm.DB
}

func Newuser(db *gorm.DB) Iuserrepo {
  return &usersrepo{
    db: db,
  }
}

func (r *usersrepo) CreateUser(user models.Users) error {
  return r.db.Create(&user).Error
}

func (r *usersrepo) GetAllUsers() ([]models.Users, error) {
  var users []models.Users

  err := r.db.Find(&users).Error

  return users, err
}

func (r *usersrepo) GetByUserID(id uint) (models.Users, error) {
  var user models.Users

  err := r.db.First(&user, id).Error

  return user, err
}

func (r *usersrepo) DeleteUser(id uint) error {
  return r.db.Delete(&models.Users{}, id).Error
}

func (r *usersrepo) UpdateUsers(id uint) error {
err := r.db.Model(&models.Users{}).Where("id = ?", id).Update("status", "done").Error

  if err != nil {
    return fmt.Errorf("r.db.Model: %w", err)
  }

  return nil
}