package auth_service

import (
	"encoding/json"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
	userCodec "github.com/ruybrito106/ads-manager-services/back-end/src/users/json"
)

func (s server) loginUserHandler(w http.ResponseWriter, r *http.Request) {

	user := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	logged, err := s.Svc.LoginUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := userCodec.UserToJSON(logged)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	user := &users.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	logged, err := s.Svc.RegisterUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := userCodec.UserToJSON(logged)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}
