package main

import (
	"fmt"
	"log"
	"net/http"

	"nd/handlers"
	"nd/internal/config"
	"nd/internal/models"
	"nd/internal/repository"
	"nd/internal/service"
	"nd/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "go.uber.org/zap"
)

func main() {
	logger, err := logger.New(true)
	if err != nil {
		log.Fatal("Failed to create logger", err)
	}


  cfg, err := config.New("./config/config.env")
	if err != nil{
		log.Fatal("config.New",err)
	}
  
  dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	

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
	userHandler := handler.New(logger, userService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users", userHandler.GetAll)

	mux.HandleFunc("GET /user", userHandler.GetByID)
	mux.HandleFunc("DELETE /user", userHandler.Delete)
	mux.HandleFunc("PATCH /user/{id}/status", userHandler.Update)

	log.Println("server started on :8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
