package auth_service

import (
	"log"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
)

const (
	entryPoint        = "/users"
	loginUserRoute    = entryPoint + "/login"
	registerUserRoute = entryPoint + "/register"
)

type server struct {
	Addr string
	Svc  Service
}

type AuthServer interface {
	ListenAndServe()
}

func NewAuthServer(addr string, logger log.Logger, iBalance balance_interface.BalanceInterface) AuthServer {
	var s AuthServer
	s = server{
		addr,
		NewService(logger, iBalance),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(loginUserRoute, s.loginUserHandler)
	http.HandleFunc(registerUserRoute, s.registerUserHandler)

	http.ListenAndServe(s.Addr, nil)
}
