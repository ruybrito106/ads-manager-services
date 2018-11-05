package balance_service

import (
	"log"
	"net/http"
)

const (
	entryPoint = "/balances"
)

type server struct {
	Addr string
	Svc  Service
}

type BalanceServer interface {
	ListenAndServe()
}

func NewBalanceServer(addr string, logger log.Logger) BalanceServer {
	var s BalanceServer
	s = server{
		addr,
		NewService(logger),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(entryPoint, s.getBalancesHandler)

	http.ListenAndServe(s.Addr, nil)
}
