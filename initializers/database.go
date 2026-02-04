package initializers

import (
	"fmt"
	"log"
	"os"
	"github.com/Hisyam/freepass-2026/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	
	dsn := os.Getenv("DB_URL")
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database:", err)
	}
	
	fmt.Println("Connected to Database successfully!")
}

func SyncDatabase() {
	modelsToMigrate := []interface{}{
		&models.User{},
		&models.Canteen{},
		&models.MenuItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.Feedback{},
	}

	for _, model := range modelsToMigrate {
		err := DB.AutoMigrate(model)
		if err != nil {
			fmt.Printf("Failed to migrate model: %T, error: %v\n", model, err)
		} else {
			fmt.Printf("Migrate model successfully: %T\n", model)
		}
	}
}