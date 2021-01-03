package scrapper

import (
	"github.com/PuerkitoBio/goquery"
	"grpcChassis/auto_scout_scrapper/models"
	"log"
	"net/http"
	"strconv"
)

func Scrape(page int) ([]*models.AutoScoutCarResponse, error) {
	// Request the HTML page.
	requestString := "https://www.autoscout24.com/lst?&sort=standard&desc=0&" +
		"offer=J%2CU%2CO%2CD%2CS&ustate=N%2CU&size=20&atype=C&page=" + strconv.Itoa(page)
	res, err := http.Get(requestString)
	if err != nil {
		log.Fatal(err)
	}
	//TODO check error
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return ParsePage(res)
}

func ParsePage(res *http.Response) ([]*models.AutoScoutCarResponse, error) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var carModels []*models.AutoScoutCarResponse

	doc.Find(".cl-list-element").Each(func(i int, s *goquery.Selection) {
		carModel := &models.AutoScoutCarResponse{}

		UUID, isUUID := s.Attr("data-guid")

		carModel.UUID = UUID
		carModel.MakeModel = s.Find(".cldt-summary-makemodel").First().Text()
		carModel.Version = s.Find(".cldt-summary-version").First().Text()
		carModel.Price = s.Find(".cldt-price").First().Text()


		// Link
		carModel.Link, _ = s.Find("a[data-item-name='detail-page-link']").First().Attr("href")

		s.Find("ul[data-item-name='vehicle-details'] li").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				carModel.Mileage = s.Text()
			case 1:
				carModel.FirstRegistration = s.Text()
			case 2:
				carModel.Power = s.Text()
			case 3:
				carModel.OfferType = s.Text()
			case 4:
				carModel.PreviousOwners = s.Text()
			case 5:
				carModel.TransmissionType = s.Text()
			case 6:
				carModel.FuelType = s.Text()
			case 7:
				carModel.FuelConsumption = s.Text()
			case 8:
				carModel.Co2Emission = s.Text()

			}
		})

		if !isUUID {
			log.Print("No uuid: %v\n", carModel)
			return
		}

		carModels = append(carModels, carModel)
	})

	return carModels, nil
}