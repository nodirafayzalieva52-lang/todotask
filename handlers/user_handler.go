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

  type UserHandler struct {
    logger      *logger.Logger
    service service.IUserService
  }

  func NewUser(logger *logger.Logger, service service.IUserService) *UserHandler {
    return &UserHandler{
      logger:    logger,
      service: service,
    }
  }

  func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.Users

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
      h.logger.Error("h.userService.Create", zap.Error(err))
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
    }

    err = h.service.CreateUser(user)
    if err != nil {
      h.logger.Error("h.userService.Create", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusCreated)
  }

  func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := h.service.GetAllUsers()
    if err != nil {
      h.logger.Error("h.userService.GetAll", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")

    err = json.NewEncoder(w).Encode(users)
    if err != nil {
      h.logger.Error("h.userService.GetAllUsers", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  func (h *UserHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
      h.logger.Error("h.userService.GetByUserId", zap.Error(err))
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    user, err := h.service.GetByUserID(uint(id))
    if err != nil {
      h.logger.Error("h.userService.GetByUserId", zap.Error(err))
      http.Error(w, err.Error(), http.StatusNotFound)
      return
    }

    w.Header().Set("Content-Type", "application/json")

    err = json.NewEncoder(w).Encode(user)
    if err != nil {
      h.logger.Error("h.userService.GetById", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
      h.logger.Error("h.userService.Delete", zap.Error(err))
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    err = h.service.DeleteUser(uint(id))
    if err != nil {
      h.logger.Error("h.userService.Delete", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusNoContent)
  }

  func (h *UserHandler) UpdateUsers(w http.ResponseWriter, r *http.Request) {
      idStr := r.PathValue("id")
      id, err := strconv.Atoi(idStr)
    if err != nil {
      h.logger.Error("h.userService.Delete", zap.Error(err))
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    err = h.service.UpdateByUserId(uint(id))
    if err != nil {
      h.logger.Error("h.userService.Update", zap.Error(err))
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.WriteHeader(http.StatusNoContent)
  }