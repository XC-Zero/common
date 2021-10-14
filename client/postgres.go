package client

import (
	"fmt"

	cfg "common/config"
	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitPostgres returns a Postgres DB engine from config
func InitPostgres(config cfg.PostgresConfiguration) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.DBName,
		config.Password,
	)

	db, err := gorm.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)
	db.LogMode(config.LogMode)

	return db, nil
}
