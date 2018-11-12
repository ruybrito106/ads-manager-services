package place_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
)

type PlaceInterface interface {
	CreatePlace(*places.Place) (*places.Place, error)
	DeletePlace(string) error
	GetPlaces() ([]*places.Place, error)
}
