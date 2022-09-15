package orm

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	host := "localhost"
	user := "zazh"
	password := "1edho222"
	dbName := "rentaldb"

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, errors.New("gagal init database")
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)

	return gormDB, nil
}
