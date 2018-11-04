package gateway_service

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

type Service interface {
	// Campaign Service Methods
	GetCampaigns() ([]*campaigns.Campaign, error)
	CreateCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	EditCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
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


func (s basicService) GetCampaigns() ([]*campaigns.Campaign, error) {
	return s.ICampaign.GetCampaigns()
}

func (s basicService) EditCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaign.EditCampaign(campaign)
}

func (s basicService) CreateCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaign.CreateCampaign(campaign)
}

func (s basicService) PauseCampaign(id int32) error {
	return s.ICampaign.PauseCampaign(id)
}
