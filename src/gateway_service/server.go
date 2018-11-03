package gateway_service

import (
	"net/http"

	"github.com/ruybrito106/ads-manager-services/src/campaign_interface"
)

const entryPoint = "/"

const (
	campaignEntryPoint  = entryPoint + "campaigns"
	createCampaignRoute = campaignEntryPoint + "/create"
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

	http.HandleFunc(createCampaignRoute, s.createCampaignHandler)
	http.HandleFunc(pauseCampaignRoute, s.pauseCampaignHandler)

	http.ListenAndServe(s.Addr, nil)
}
