package service

import (
	"errors"
	"nd/internal/models"
	"nd/internal/repository"
	"strings"
)

type IUserService interface {
	CreateUser(user models.Users) error
	GetAllUsers() ([]models.Users, error)
	GetByUserID(id uint) (models.Users, error)
	DeleteUser(id uint) error
	UpdateByUserId(id uint) error
	
}

type userservice struct {
	usersrepo repository.Iuserrepo
}

func NewUser(usersrepo repository.Iuserrepo) IUserService {
	return &userservice{
		usersrepo: usersrepo,
  }
}

func (s *userservice) CreateUser(user models.Users) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("empty name")
	}

	return s.usersrepo.CreateUser(user)
}

func (s *userservice) GetAllUsers() ([]models.Users, error) {
	return s.usersrepo.GetAllUsers()
}

func (s *userservice) GetByUserID(id uint) (models.Users, error) {
	if id == 0 {
		return models.Users{}, errors.New("invalid id")
	}
	return s.usersrepo.GetByUserID(id)
}

func (s *userservice) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.usersrepo.DeleteUser(id)
}

func (s *userservice) UpdateByUserId(id uint) error{
	return s.usersrepo.UpdateUsers(id)
}