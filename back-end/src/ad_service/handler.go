package ad_service

import (
	"encoding/json"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
	adCodec "github.com/ruybrito106/ads-manager-services/back-end/src/ads/json"
)

func (s server) createAdHandler(w http.ResponseWriter, r *http.Request) {

	ad := &ads.Ad{}
	if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	created, err := s.Svc.CreateAd(ad)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := adCodec.AdToJSON(created)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}

func (s server) deleteAdHandler(w http.ResponseWriter, r *http.Request) {

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := s.Svc.DeleteAd(idStrs[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s server) getAdsHandler(w http.ResponseWriter, r *http.Request) {

	var idStrs []string
	if idStrs = r.URL.Query()["id"]; len(idStrs) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	ads, err := s.Svc.GetAds(idStrs[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoded, err := adCodec.AdsToJSON(ads)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(encoded)
}
