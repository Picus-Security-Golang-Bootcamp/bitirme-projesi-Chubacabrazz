package db

import (
	"fmt"
	"os"
	"time"

	"github.com/Chubacabrazz/picus-storeApp/pkg/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Func Connect: Trying to connect to the database and initial settings, returns fatal error if can't reach to db.
func Connect(cfg *config.Config) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		zap.L().Fatal("Cannot load env file", zap.Error(err))
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("PICUS_DB_HOST"),
		os.Getenv("PICUS_DB_PORT"),
		os.Getenv("PICUS_DB_USERNAME"),
		os.Getenv("PICUS_DB_NAME"),
		os.Getenv("PICUS_DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot connect to database", zap.Error(err))
	}

	origin, err := db.DB()
	if err != nil {
		zap.L().Fatal("Cannot get info from database", zap.Error(err))
	}

	err = origin.Ping()
	if err != nil {
		zap.L().Fatal("Cannot ping database", zap.Error(err))
	}

	origin.SetMaxOpenConns(cfg.DBConfig.MaxOpen)
	origin.SetMaxIdleConns(cfg.DBConfig.MaxIdle)
	origin.SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxLifetime) * time.Second)

	return db
}
