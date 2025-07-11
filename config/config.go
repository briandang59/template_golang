package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/briandang59/be_scada/internal/model"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() {
	_ = godotenv.Load() // đọc .env

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("🚫 Không kết nối được DB: %v", err)
	}

	DB = db
	log.Println("✅ Đã kết nối PostgreSQL")

	// Auto‑migrate tất cả model cần thiết
	autoMigrate()
}

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("🚫 JWT_SECRET không được để trống trong .env")
	}
	return secret
}

func autoMigrate() {
	if err := DB.AutoMigrate(
		&model.Factory{},
		&model.Department{},
		&model.EquipmentType{},
		&model.Equipment{},
		&model.Account{},
	); err != nil {
		log.Fatalf("🚫 AutoMigrate lỗi: %v", err)
	}
}
