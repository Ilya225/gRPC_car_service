package models

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"strconv"
	"strings"
	"time"
)

type AutoScoutCarResponse struct {
	UUID              string
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
	UUID                 string        `db:"uuid"`
	Source               string        `db:"source"`
	Link                 string        `db:"link"`
	Make                 string        `db:"make"`
	Model                string        `db:"model"`
	Version              string        `db:"version"`
	Price                *sql.NullInt32 `db:"price"`
	Currency             string        `db:"currency"`
	Mileage              *sql.NullInt32 `db:"mileage"`
	MileageUnit          string        `db:"mileage_unit"`
	FirstRegistration    *sql.NullInt64 `db:"first_registration"`
	Power                *sql.NullInt32 `db:"power"`
	PowerUnit            string        `db:"power_unit"`
	OfferType            string        `db:"offer_type"`
	PreviousOwnersNumber *sql.NullInt32        `db:"previous_owners_number"`
	TransmissionType     string        `db:"transmission_type"`
	FuelType             string        `db:"fuel_type"`
	FuelConsumption      *sql.NullInt32 `db:"fuel_consumption"`
	Co2Emission          *sql.NullInt32 `db:"co2_emission"`
}

func AutoScoutResponseToDatabaseModel(response *AutoScoutCarResponse) *Car {
	var resError error

	makeModel := strings.Fields(response.MakeModel)

	priceCurrency := strings.Fields(response.Price)

	price, err := strconv.Atoi(priceCurrency[1][:len(priceCurrency[1])-2])
	priceField := &sql.NullInt32{}
	resError = multierror.Append(resError, err)
	if err == nil {
		priceField.Int32 = int32(price)
		priceField.Valid = true
	}

	mileageLine := strings.Fields(response.Mileage)
	mileage, err := strconv.Atoi(mileageLine[0])
	mileageField := &sql.NullInt32{}
	resError = multierror.Append(resError, err)
	if err == nil {
		mileageField.Int32 = int32(mileage)
		mileageField.Valid = true
	}

	layout := "01/2006"
	firstRegistration, err := time.Parse(layout, response.FirstRegistration)
	firstRegisField := &sql.NullInt64{}
	resError = multierror.Append(resError, err)
	if err == nil {
		firstRegisField.Int64 = firstRegistration.Unix()
		firstRegisField.Valid = true
	}

	powerLine := strings.Replace(response.Power, "(", "", -1)
	powerLine = strings.Replace(powerLine, ")", "", -1)
	powerSlice := strings.Fields(powerLine)
	power, err := strconv.Atoi(powerSlice[0])
	powerField := &sql.NullInt32{}
	resError = multierror.Append(resError, err)
	if err == nil {
		powerField.Int32 = int32(power)
		powerField.Valid = true
	}

	previousOwners := strings.Fields(response.PreviousOwners)
	prevOwnerNum, err := strconv.Atoi(previousOwners[0])
	prevOwnerNumField := &sql.NullInt32{}
	resError = multierror.Append(resError, err)
	if err == nil {
		prevOwnerNumField.Int32 = int32(prevOwnerNum)
		prevOwnerNumField.Valid = true
	}

	return &Car{
		UUID:              response.UUID,
		Source:            "AutoScout24",
		Link:              response.Link,
		Make:              makeModel[0],
		Model:             makeModel[1],
		Version:           response.Version,
		Price:             priceField,
		Currency:          priceCurrency[0],
		Mileage:           mileageField,
		MileageUnit:       mileageLine[0],
		FirstRegistration: firstRegisField,
		Power:             powerField,
		PowerUnit:         powerSlice[1],
		OfferType:         response.OfferType,
		PreviousOwnersNumber: prevOwnerNumField,
		TransmissionType:  response.TransmissionType,
		FuelType:          response.FuelType,
		//FuelConsumption:      0,
		//Co2Emission:          0,
	}
}
