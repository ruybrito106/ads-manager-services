package json

import (
	"encoding/json"

	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
)

func PlaceToJSON(place *places.Place) ([]byte, error) {
	return json.Marshal(&place)
}

func PlaceFromJSON(encoded []byte) (*places.Place, error) {
	place := places.Place{}
	if err := json.Unmarshal(encoded, &place); err != nil {
		return nil, err
	}
	return &place, nil
}

func PlacesToJSON(places []*places.Place) ([]byte, error) {
	return json.Marshal(&places)
}

func PlacesFromJSON(encoded []byte) ([]*places.Place, error) {
	places := []*places.Place{}
	if err := json.Unmarshal(encoded, &places); err != nil {
		return nil, err
	}
	return places, nil
}
