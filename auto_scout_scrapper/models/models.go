package models

type AutoScoutCarResponse struct {
	Link              string
	MakeModel         string
	Version           string
	Price             string
	Mileage           string
	FirstRegistration string
	Power             string
	OfferType         string
	PreviousOwners    string
	TransmissionType  string
	FuelType          string
	FuelConsumption   string
	Co2Emission       string
}

type Car struct {
	UUID                 string
	Source               string
	Link                 string
	MakeModel            string
	Version              string
	Price                int
	Currency             string
	Mileage              int
	MileageUnit          string
	FirstRegistration    string
	Power                int
	PowerUnit            string
	OfferType            string
	PreviousOwnersNumber string
	TransmissionType     string
	FuelType             string
	FuelConsumption      int
	Co2Emission          int
}
