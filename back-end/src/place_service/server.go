package place_service

import (
	"log"
	"net/http"
)

const (
	entryPoint       = "/places"
	createPlaceRoute = entryPoint + "/create"
	deletePlaceRoute = entryPoint + "/delete"
)

type server struct {
	Addr string
	Svc  Service
}

type PlaceServer interface {
	ListenAndServe()
}

func NewPlaceServer(addr string, logger log.Logger) PlaceServer {
	var s PlaceServer
	s = server{
		addr,
		NewService(logger),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(entryPoint, s.getPlacesHandler)
	http.HandleFunc(deletePlaceRoute, s.deletePlaceHandler)
	http.HandleFunc(createPlaceRoute, s.createPlaceHandler)

	http.ListenAndServe(s.Addr, nil)
}
