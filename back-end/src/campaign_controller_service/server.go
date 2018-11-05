package campaign_controller_service

import (
	"log"
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_interface"
)

const (
	entryPoint          = "/controller/campaigns"
	createCampaignRoute = entryPoint + "/create"
	editCampaignRoute   = entryPoint + "/edit"
	pauseCampaignRoute  = entryPoint + "/pause"
)

type server struct {
	Addr string
	Svc  Service
}

type CampaignControllerServer interface {
	ListenAndServe()
}

func NewCampaignControllerServer(addr string, logger log.Logger, iCampaign campaign_interface.CampaignInterface, iBalance balance_interface.BalanceInterface) CampaignControllerServer {
	var s CampaignControllerServer
	s = server{
		addr,
		NewService(
			logger,
			iCampaign,
			iBalance,
		),
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
