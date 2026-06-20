package service

import (
	"errors"
	"nd/internal/models"
	"nd/internal/repository"
	"strings"
)

type ItaskService interface {
	Create(task models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id uint) (models.Task, error)
	Delete(id uint) error
	UpdateById(id uint) error
	
}

type service struct {
	repo repository.ItaskRepo
	userRepo repository.Iuserrepo
}

func New(repo repository.ItaskRepo) ItaskService {
	return &service{
		repo: repo,
  }
}

func (s *service) Create(task models.Task) error {
	if strings.TrimSpace(task.Title) == "" {
		return errors.New("empty title")
	}

	return s.repo.Create(task)
}

func (s *service) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (models.Task, error) {
	if id == 0 {
		return models.Task{}, errors.New("invalid id")
	}
	return s.repo.GetByID(id)
}

func (s *service) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(id)
}

func (s *service) UpdateById(id uint) error{
	return s.repo.Update(id)
}