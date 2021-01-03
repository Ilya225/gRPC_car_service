package car_storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"grpcChassis/auto_scout_scrapper/models"
	"io"
)

const (
	carInsert = "INSERT INTO car " +
		"(source_id, source, link, make, model, version, price, currency, mileage, mileage_unit, first_registration, " +
		"power, power_unit, offer_type, previous_owners_number, transmission_type, fuel_type, fuel_consumption, co2_emission) " +
		"VALUES (:source_id, :source, :link, :make, :model, :version, :price, :currency, :mileage, :mileage_unit, :first_registration, " +
		":power, :power_unit, :offer_type, :previous_owners_number, " +
		":transmission_type, :fuel_type, :fuel_consumption, :co2_emission) ON CONFLICT DO NOTHING"
)

type Repository struct {
	io.Closer
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveCar(car *models.Car) error {
	tx := r.db.MustBegin()
	_, err := tx.NamedExec(carInsert, car)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (r *Repository) Close() error {
	return r.db.Close()
}
