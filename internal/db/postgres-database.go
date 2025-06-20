package db

import (
	"log"

	"github.com/joaopedro489/back-golang/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	db *gorm.DB
}

func NewPostgresDatabase() *PostgresDatabase {
	db, err := gorm.Open(postgres.Open(config.Env.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	return &PostgresDatabase{db: db}
}

func (p *PostgresDatabase) GetDB() *gorm.DB {
	if p.db == nil {
		p = NewPostgresDatabase()
	}
	return p.db
}
