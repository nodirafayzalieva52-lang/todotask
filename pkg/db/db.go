package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres" 
	"gorm.io/gorm"          
)

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type DB struct {
	GormDB *gorm.DB
}

func New(opts Options) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		opts.Host, opts.User, opts.Password, opts.DBName, opts.Port)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(4)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return &DB{GormDB: gormDB}, nil
}

func (d *DB) Close() {
	sqlDB, err := d.GormDB.DB()
	if err == nil {
		sqlDB.Close()
	}
}