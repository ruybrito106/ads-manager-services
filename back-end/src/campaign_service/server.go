package campaign_service

import (
	"log"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
)

const (
	entryPoint          = "/campaigns"
	createCampaignRoute = entryPoint + "/create"
	editCampaignRoute   = entryPoint + "/edit"
	pauseCampaignRoute  = entryPoint + "/pause"
)

type server struct {
	Addr string
	Svc  Service
}

type CampaignServer interface {
	ListenAndServe()
}

func NewCampaignServer(addr string, logger log.Logger, iBalance balance_interface.BalanceInterface) CampaignServer {
	var s CampaignServer
	s = server{
		addr,
		NewService(logger, iBalance),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(entryPoint, s.getCampaignsHandler)
	http.HandleFunc(editCampaignRoute, s.editCampaignHandler)
	http.HandleFunc(createCampaignRoute, s.createCampaignHandler)
	http.HandleFunc(pauseCampaignRoute, s.pauseCampaignHandler)

	http.ListenAndServe(s.Addr, nil)
}
