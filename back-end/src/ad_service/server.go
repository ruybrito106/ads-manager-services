package ad_service

import (
	"log"
	"net/http"
)

const (
	entryPoint    = "/ads"
	createAdRoute = entryPoint + "/create"
	deleteAdRoute = entryPoint + "/delete"
)

type server struct {
	Addr string
	Svc  Service
}

type AdServer interface {
	ListenAndServe()
}

func NewAdServer(addr string, logger log.Logger) AdServer {
	var s AdServer
	s = server{
		addr,
		NewService(logger),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(entryPoint, s.getAdsHandler)
	http.HandleFunc(deleteAdRoute, s.deleteAdHandler)
	http.HandleFunc(createAdRoute, s.createAdHandler)

	http.ListenAndServe(s.Addr, nil)
}
