package campaign_service

import (
	"log"
	"net/http"
)

const (
	entryPoint          = "/campaigns"
	createCampaignRoute = entryPoint + "/create"
	pauseCampaignRoute  = entryPoint + "/pause"
)

type server struct {
	Addr string
	Svc  Service
}

type CampaignServer interface {
	ListenAndServe()
}

func NewCampaignServer(addr string, logger log.Logger) CampaignServer {
	var s CampaignServer
	s = server{
		addr,
		NewService(logger),
	}
	return s
}

func (s server) ListenAndServe() {

	http.HandleFunc(entryPoint, s.getCampaignsHandler)
	http.HandleFunc(createCampaignRoute, s.createCampaignHandler)
	http.HandleFunc(pauseCampaignRoute, s.pauseCampaignHandler)

	http.ListenAndServe(s.Addr, nil)
}
