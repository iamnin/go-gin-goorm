package config

import (
	"fmt"
	"go-gin-goorm/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "172.20.0.2"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbName   = "goginorm"
)

// Connect to database
func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
