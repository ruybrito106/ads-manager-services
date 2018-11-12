package place_service

import (
	"encoding/json"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
	placeCodec "github.com/ruybrito106/ads-manager-services/back-end/src/places/json"
)

func (s server) createPlaceHandler(w http.ResponseWriter, r *http.Request) {

	place := &places.Place{}
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	created, err := s.Svc.CreatePlace(place)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := placeCodec.PlaceToJSON(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) deletePlaceHandler(w http.ResponseWriter, r *http.Request) {

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := s.Svc.DeletePlace(idStrs[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s server) getPlacesHandler(w http.ResponseWriter, r *http.Request) {

	places, err := s.Svc.GetPlaces()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := placeCodec.PlacesToJSON(places)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}
