package database

import (
	"github.com/hikaru-shindo/fiber-playground/internal/data"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGormSqliteDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn))

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func GormTestSqliteInMemoryDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:?cache=shared"))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GormMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&data.Product{},
	)
}
