package database

import (
	"database/sql"
	"os"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func NewBunSqliteDatabase(dsn string) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dsn)

	if err != nil {
		return nil, err
	}

	sqldb.SetMaxIdleConns(3)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Hour)

	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}

func BunTestSqliteDatabase(dsn string) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dsn)

	if err != nil {
		return nil, err
	}

	return bun.NewDB(sqldb, sqlitedialect.New()), nil
}

func BunDropTestSqliteDatabase(dsn string) error {
	if err := os.Remove(dsn); err != nil {
		return err
	}
	return nil
}

func BunMigrate(db *bun.DB) error {
	// todo
	return nil
}
