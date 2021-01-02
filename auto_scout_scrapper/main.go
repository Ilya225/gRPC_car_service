package main

import (
	"fmt"
	"grpcChassis/auto_scout_scrapper/scrapper"
	"log"
)

func main() {
	scrape, err := scrapper.Scrape()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("%v\n", scrape)
}
