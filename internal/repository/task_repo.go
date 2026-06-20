package repository

import (
	"fmt"
	"nd/internal/models"

	"gorm.io/gorm"
)

type ItaskRepo interface {
  Create(task models.Task) error
  GetAll() ([]models.Task, error)
  GetByID(id uint) (models.Task, error)
  Delete(id uint) error
  Update(id uint) error
}

type repo struct {
  db *gorm.DB
}

func New(db *gorm.DB) ItaskRepo {
  return &repo{
    db: db,
  }
}

func (r *repo) Create(task models.Task) error {
  return r.db.Create(&task).Error
}

func (r *repo) GetAll() ([]models.Task, error) {
  var tasks []models.Task

  err := r.db.Find(&tasks).Error

  return tasks, err
}

func (r *repo) GetByID(id uint) (models.Task, error) {
  var task models.Task

  err := r.db.First(&task, id).Error

  return task, err
}

func (r *repo) Delete(id uint) error {
  return r.db.Delete(&models.Task{}, id).Error
}

func (r *repo) Update(id uint) error {
err := r.db.Model(&models.Task{}).Where("id = ?", id).Update("status", "done").Error

  if err != nil {
    return fmt.Errorf("r.db.Model: %w", err)
  }

  return nil
}