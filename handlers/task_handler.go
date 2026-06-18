package handler

import (
  "encoding/json"
  "strconv"
  "net/http"
  "nd/internal/models"
  "nd/internal/service"
)

type TaskHandler struct {
  service service.ItaskService
}

func New(service service.ItaskService) TaskHandler {
  return TaskHandler{
    service: service,
  }
}

func (h TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
  var task models.Task

  err := json.NewDecoder(r.Body).Decode(&task)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  err = h.service.Create(task)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
}

func (h TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
  tasks, err := h.service.GetAll()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")

  err = json.NewEncoder(w).Encode(tasks)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func (h TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
  idStr := r.URL.Query().Get("id")

  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "invalid id", http.StatusBadRequest)
    return
  }

  task, err := h.service.GetByID(uint(id))
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }

  w.Header().Set("Content-Type", "application/json")

  err = json.NewEncoder(w).Encode(task)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func (h TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
  idStr := r.URL.Query().Get("id")

  id, err := strconv.Atoi(idStr)
  if err != nil {
    http.Error(w, "invalid id", http.StatusBadRequest)
    return
  }

  err = h.service.Delete(uint(id))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusNoContent)
}