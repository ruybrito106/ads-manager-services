package gateway_service

import (
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_interface"
)

const entryPoint = "/"

const (
	campaignEntryPoint  = entryPoint + "campaigns"
	createCampaignRoute = campaignEntryPoint + "/create"
	editCampaignRoute = campaignEntryPoint + "/edit"
	pauseCampaignRoute  = campaignEntryPoint + "/pause"

	authEntryPoint    = entryPoint + "auth"
	registerUserRoute = authEntryPoint + "/register"
	loginUserRoute    = authEntryPoint + "/login"
)

type server struct {
	Addr string
	Svc  Service
}

type GatewayServer interface {
	ListenAndServe()
}

func NewGatewayServer(addr string, iCampaign campaign_interface.CampaignInterface) GatewayServer {
	var s GatewayServer
	s = server{
		addr,
		iCampaign,
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(campaignEntryPoint, s.getCampaignsHandler)
	http.HandleFunc(createCampaignRoute, s.createCampaignHandler)
	http.HandleFunc(editCampaignRoute, s.editCampaignHandler)
	http.HandleFunc(pauseCampaignRoute, s.pauseCampaignHandler)

	http.ListenAndServe(s.Addr, nil)
}
