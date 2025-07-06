package database

import (
	"fmt"

	"github.com/gsc23/jemax-bot/pkg/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type Config struct {
    Host     string
    Port     int
    Database string
    Username string
    Password string
}

func NewDatabase(cfg Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port,
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logs.DefaultLoggerWithName("database"),
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	
	return &Database{
		DB: db,
	}, nil
}