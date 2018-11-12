package store

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/place_service/store/postgres"
	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
)

type Store interface {
	CreatePlace(*places.Place) (*places.Place, error)
	DeletePlace(string) error
	GetPlaces() ([]*places.Place, error)
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) basicStore {
	return basicStore{logger}
}

func (s basicStore) CreatePlace(place *places.Place) (*places.Place, error) {
	db := postgres.NewDatabase(s.logger)
	return db.CreatePlace(place)
}

func (s basicStore) DeletePlace(id string) error {
	db := postgres.NewDatabase(s.logger)
	return db.DeletePlace(id)
}

func (s basicStore) GetPlaces() ([]*places.Place, error) {
	db := postgres.NewDatabase(s.logger)
	return db.GetPlaces()
}
