package main

import (
  "log"
  "net/http"

  "nd/handlers"
  "nd/internal/models"
  "nd/internal/repository"
  "nd/internal/service"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func main() {
  dsn := "host=localhost user=postgres password=20102010 dbname=task port=5432 sslmode=disable"

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
  }

  err = db.AutoMigrate(&models.Task{})
  if err != nil {
    log.Fatal(err)
  }

  userRepo := repository.New(db)
  userService := service.New(userRepo)
  userHandler := handler.New(userService)

  mux := http.NewServeMux()

  mux.HandleFunc("POST /users", userHandler.Create)
  mux.HandleFunc("GET /users", userHandler.GetAll)

  mux.HandleFunc("GET /user", userHandler.GetByID)
  mux.HandleFunc("DELETE /user", userHandler.Delete)

  log.Println("server started on :8080")

  err = http.ListenAndServe(":8080", mux)
  if err != nil {
    log.Fatal(err)
  }
}