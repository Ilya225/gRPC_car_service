package car_storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"grpcChassis/auto_scout_scrapper/env"
	"log"
)

func Connect(postgreConfig* env.PostgreConfig) *sqlx.DB {
	db, err := sqlx.Connect("postgres", postgreConfig.DataSourceName())
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
