package store

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_service/store/postgres"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

type Store interface {
	EditCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	CreateCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	GetCampaigns() ([]*campaigns.Campaign, error)

	PauseCampaign(id int32) error
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) basicStore {
	return basicStore{logger}
}

func (s basicStore) EditCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	db := postgres.NewDatabase(s.logger)
	return db.EditCampaign(campaign)
}

func (s basicStore) CreateCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	db := postgres.NewDatabase(s.logger)
	return db.CreateCampaign(campaign)
}

func (s basicStore) GetCampaigns() ([]*campaigns.Campaign, error) {
	db := postgres.NewDatabase(s.logger)
	return db.GetCampaigns()
}

func (s basicStore) PauseCampaign(id int32) error {
	db := postgres.NewDatabase(s.logger)
	return db.PauseCampaign(id)
}
