package balance_service

import (
	"net/http"

	balanceCodec "github.com/ruybrito106/ads-manager-services/back-end/src/balances/json"
)

func (s server) getBalancesHandler(w http.ResponseWriter, r *http.Request) {

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	balance, err := s.Svc.GetBalanceByUserID(idStrs[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := balanceCodec.BalanceToJSON(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) initBalanceHandler(w http.ResponseWriter, r *http.Request) {

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	balance, err := s.Svc.InitBalance(idStrs[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := balanceCodec.BalanceToJSON(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}
