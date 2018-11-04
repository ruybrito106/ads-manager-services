package campaign_service

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_service/store"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

type Service interface {
	GetCampaigns() ([]*campaigns.Campaign, error)
	CreateCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	EditCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error
}

type basicService struct {
	store store.Store
}

func NewService(logger log.Logger) Service {
	var service Service
	service = basicService{store.New(logger)}
	return service
}

func (s basicService) GetCampaigns() ([]*campaigns.Campaign, error) {
	return s.store.GetCampaigns()
}

func (s basicService) EditCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.store.EditCampaign(campaign)
}

func (s basicService) CreateCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.store.CreateCampaign(campaign)
}

func (s basicService) PauseCampaign(id int32) error {
	return s.store.PauseCampaign(id)
}