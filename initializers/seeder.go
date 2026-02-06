package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/Hisyam/freepass-2026/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var count int64
	if err := DB.Model(&models.User{}).Where("role = ?", "ADMIN").Count(&count).Error; err != nil {
		log.Printf("Error checking admin count: %v", err)
		return
	}

	if count > 0 {
		fmt.Println("Admin already exists, skipping seeder.")
		return
	}

	fmt.Println("No Admin found. Seeding Super Admin...")

	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		password = "admin123"
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	admin := models.User{
		ID:       uuid.New(),
		Name:     "Super Admin",
		Email:    "admin@freepass.com",
		Password: string(hashedPassword),
		Role:     "ADMIN",
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Fatalf("Failed to seed admin: %v", err)
	}

	fmt.Println("Super Admin created successfully")
}