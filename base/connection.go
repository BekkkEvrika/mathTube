package base

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Connect ...
func Connect() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", "postgres", "postgres", "math_tube", 5433)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return err
	}
	return nil
}

func Migrate(v interface{}) error {
	return db.Migrator().AutoMigrate(v)
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
