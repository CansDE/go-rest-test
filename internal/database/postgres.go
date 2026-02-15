package database

import (
	"fmt"
	"test-rest-api/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const maxLifetime = 5 * time.Minute

func NewPostgres(config config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(maxLifetime)
	sqlDB.SetConnMaxIdleTime(maxLifetime)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return db, nil
}
