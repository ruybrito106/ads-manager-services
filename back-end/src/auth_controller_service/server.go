package auth_controller_service

import (
	"log"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
)

const (
	entryPoint        = "/controller/users"
	loginUserRoute    = entryPoint + "/login"
	registerUserRoute = entryPoint + "/register"
)

type server struct {
	Addr string
	Svc  Service
}

type AuthControllerServer interface {
	ListenAndServe()
}

func NewAuthControllerServer(addr string, logger log.Logger, iAuth auth_interface.AuthInterface, iBalance balance_interface.BalanceInterface) AuthControllerServer {
	var s AuthControllerServer
	s = server{
		addr,
		NewService(
			logger,
			iAuth,
			iBalance,
		),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(loginUserRoute, s.loginUserHandler)
	http.HandleFunc(registerUserRoute, s.registerUserHandler)

	http.ListenAndServe(s.Addr, nil)
}
