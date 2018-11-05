package gateway_service

import (
	"net/http"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_controller_interface"
)

const entryPoint = "/"

const (
	campaignEntryPoint  = entryPoint + "campaigns"
	createCampaignRoute = campaignEntryPoint + "/create"
	editCampaignRoute   = campaignEntryPoint + "/edit"
	pauseCampaignRoute  = campaignEntryPoint + "/pause"

	authEntryPoint = entryPoint + "users"
	loginUserRoute = authEntryPoint + "/login"
)

type server struct {
	Addr string
	Svc  Service
}

type GatewayServer interface {
	ListenAndServe()
}

func NewGatewayServer(addr string, iCampaignController campaign_controller_interface.CampaignControllerInterface, iAuth auth_interface.AuthInterface) GatewayServer {
	var s GatewayServer
	s = server{
		addr,
		NewService(iCampaignController, iAuth),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(campaignEntryPoint, s.getCampaignsHandler)
	http.HandleFunc(createCampaignRoute, s.createCampaignHandler)
	http.HandleFunc(editCampaignRoute, s.editCampaignHandler)
	http.HandleFunc(pauseCampaignRoute, s.pauseCampaignHandler)

	http.HandleFunc(loginUserRoute, s.loginUserHandler)

	http.ListenAndServe(s.Addr, nil)
}
