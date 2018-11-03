package gateway_service

import (
	"github.com/ruybrito106/ads-manager-services/src/campaign_interface"
	"github.com/ruybrito106/ads-manager-services/src/campaigns"
)

type Service interface {
	// Campaign Service Methods
	CreateCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error

	// Auth Service Methods
}

type basicService struct {
	ICampaign campaign_interface.CampaignInterface
}

func NewService(iCampaign campaign_interface.CampaignInterface) Service {
	var service Service
	service = basicService{
		ICampaign: iCampaign,
	}
	return service
}

func (s basicService) CreateCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaign.CreateCampaign(campaign)
}

func (s basicService) PauseCampaign(id int32) error {
	return s.ICampaign.PauseCampaign(id)
}
