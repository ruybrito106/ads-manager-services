package gateway_service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ruybrito106/ads-manager-services/back-end/src/users"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
	campaignCodec "github.com/ruybrito106/ads-manager-services/back-end/src/campaigns/json"
	userCodec "github.com/ruybrito106/ads-manager-services/back-end/src/users/json"
)

func (s server) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

func (s server) getCampaignsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	campaigns, err := s.Svc.GetCampaigns()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := campaignCodec.CampaignsToJSON(campaigns)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) createCampaignHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		idStrs = []string{"superuser"}
		// w.WriteHeader(http.StatusUnprocessableEntity)
		// return
	}

	campaign := &campaigns.Campaign{}
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	created, err := s.Svc.CreateCampaign(idStrs[0], campaign)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := campaignCodec.CampaignToJSON(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) editCampaignHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		idStrs = []string{"superuser"}
		// w.WriteHeader(http.StatusUnprocessableEntity)
		// return
	}

	campaign := &campaigns.Campaign{}
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	created, err := s.Svc.EditCampaign(idStrs[0], campaign)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := campaignCodec.CampaignToJSON(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) pauseCampaignHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var (
		id  int
		err error
	)

	if id, err = strconv.Atoi(idStrs[0]); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := s.Svc.PauseCampaign(int32(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
