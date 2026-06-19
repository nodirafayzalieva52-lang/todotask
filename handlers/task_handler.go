  package handler

  import (
    "encoding/json"
    "nd/internal/models"
    "nd/internal/service"
    "nd/pkg/logger"
    "net/http"
    "strconv"

    "go.uber.org/zap"
  )

  type TaskHandler struct {
    logger      *logger.Logger
    service service.ItaskService
  }

  func New(logger *logger.Logger, service service.ItaskService) *TaskHandler {
    return &TaskHandler{
      logger:    logger,
      service: service,
    }
  }

  func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
    var task models.Task

    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
      h.logger.Error("h.userService.Create", zap.Error(err))
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
    }

    err = h.service.Create(task)
    if err != nil {
      h.logger.Error("h.userService.Create", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusCreated)
  }

  func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.service.GetAll()
    if err != nil {
      h.logger.Error("h.userService.GetAll", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")

    err = json.NewEncoder(w).Encode(tasks)
    if err != nil {
      h.logger.Error("h.userService.GetAll", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
      h.logger.Error("h.userService.GetById", zap.Error(err))
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    task, err := h.service.GetByID(uint(id))
    if err != nil {
      h.logger.Error("h.userService.GetById", zap.Error(err))
      http.Error(w, err.Error(), http.StatusNotFound)
      return
    }

    w.Header().Set("Content-Type", "application/json")

    err = json.NewEncoder(w).Encode(task)
    if err != nil {
      h.logger.Error("h.userService.GetById", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
      h.logger.Error("h.userService.Delete", zap.Error(err))
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    err = h.service.Delete(uint(id))
    if err != nil {
      h.logger.Error("h.userService.Delete", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusNoContent)
  }