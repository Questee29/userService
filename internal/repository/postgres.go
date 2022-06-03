package repository

import (
	"TaxiApp1/config"
	"log"

	"github.com/jmoiron/sqlx"
)

type ConfigDB struct {
	DBDriver      string
	DBSource      string
	ServerAddress string
}

func NewPostgresDB(cfg ConfigDB) (*sqlx.DB, error) {
	//load env
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	//db connect
	db, err := sqlx.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
