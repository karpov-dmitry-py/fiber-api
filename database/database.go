package database

import (
	"fmt"
	"log"

	"github.com/karpov-dmitry-py/fiber-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() (*DbInstance, error) {
	var (
		err error
	)

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		err := fmt.Errorf("failed to connect to db: %v", err)
		return nil, err
	}

	log.Print("connected to db successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Print("running migrations")
	if err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}); err != nil {
		err = fmt.Errorf("failed to run db migrations: %v", err)
		return nil, err
	}

	Database = DbInstance{Db: db}

	return &Database, nil
}
