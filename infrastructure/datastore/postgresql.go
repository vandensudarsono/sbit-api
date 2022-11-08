package datastore

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func CreateConnection() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.postgres.host"), viper.GetInt32("database.postgres.port"),
		viper.GetString("database.postgres.username"), viper.GetString("database.postgres.password"),
		viper.GetString("database.postgres.dbname"))

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(10)
}

func GetConnection() *sqlx.DB {
	return db
}
