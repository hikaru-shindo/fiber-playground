package store_test

import (
	"log"
	"math/rand"

	"github.com/hikaru-shindo/fiber-playground/internal/data"
	"github.com/hikaru-shindo/fiber-playground/internal/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// newTestProduct creates a new product with test data
func newTestProduct() data.Product {
	product := new(data.Product)

	product.Id = uuid.New()
	product.Name = uuid.NewString()
	product.Description = uuid.NewString()

	product.Price.Value = rand.Int()
	product.Price.Currency = uuid.NewString()

	return *product
}

func setupGorm() *gorm.DB {
	db, err := database.GormTestSqliteDatabase("./../../database/database_test.go")

	if err != nil {
		log.Fatal(err)
	}

	if err := database.GormMigrate(db); err != nil {
		log.Fatal(err)
	}

	return db
}

func teardownGorm() {
	if err := database.GormDropTestSqliteDatabase("./../../database/database_test.go"); err != nil {
		log.Fatal(err)
	}
}
