package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345678 dbname=fourth sslmode=disable Timezone=Asia/Tashkent"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
