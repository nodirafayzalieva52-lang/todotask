package service

import (
  "nd/internal/models"
  "nd/internal/repository"
)

type ItaskService interface {
  Create(task models.Task) error
  GetAll() ([]models.Task, error)
  GetByID(id uint) (models.Task, error)
  Delete(id uint) error
}

type service struct {
  repo repository.ItaskRepo
}

func New(repo repository.ItaskRepo) ItaskService {
  return &service{
    repo: repo,
  }
}

func (s *service) Create(task models.Task) error {
  return s.repo.Create(task)
}

func (s *service) GetAll() ([]models.Task, error) {
  return s.repo.GetAll()
}

func (s *service) GetByID(id uint) (models.Task, error) {
  return s.repo.GetByID(id)
}

func (s *service) Delete(id uint) error {
  return s.repo.Delete(id)
}