package place_service

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/place_service/store"
	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
)

type Service interface {
	CreatePlace(*places.Place) (*places.Place, error)
	DeletePlace(string) error
	GetPlaces() ([]*places.Place, error)
}

type basicService struct {
	store store.Store
}

func NewService(logger log.Logger) Service {
	var service Service
	service = basicService{store.New(logger)}
	return service
}

func (s basicService) CreatePlace(place *places.Place) (*places.Place, error) {
	return s.store.CreatePlace(place)
}

func (s basicService) DeletePlace(id string) error {
	return s.store.DeletePlace(id)
}

func (s basicService) GetPlaces() ([]*places.Place, error) {
	return s.store.GetPlaces()
}
