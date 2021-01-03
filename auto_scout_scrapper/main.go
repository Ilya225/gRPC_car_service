package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"grpcChassis/auto_scout_scrapper/car_storage"
	"grpcChassis/auto_scout_scrapper/env"
	"grpcChassis/auto_scout_scrapper/models"
	"grpcChassis/auto_scout_scrapper/scrapper"
	"log"
)

func main() {
	config := env.Config()

	db, err := sqlx.Connect("postgres", config.PostgreConfig.DataSourceName())
	if err != nil {
		log.Fatalln(err)
	}

	storage := car_storage.New(db)

	response, err := scrapper.Scrape(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("%v\n", response)

	for _, carModel := range response {
		err = storage.SaveCar(models.AutoScoutResponseToDatabaseModel(carModel))
	}

	if err != nil {
		log.Fatal(err)
	}

}
