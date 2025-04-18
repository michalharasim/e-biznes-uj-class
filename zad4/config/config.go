package config

import (
	"fmt"
	"zad4/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres dbname=zad4 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Nie udało się połączyć z bazą danych")
	}

	DB = db
	fmt.Println("Połączono z bazą danych")

	db.AutoMigrate(
		// &models.Category{},
		&models.Product{},
		// &models.Supplier{},
		// &models.Manufacturer{},
		// &models.Warehouse{},
	)
}
