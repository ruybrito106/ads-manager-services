package campaign_service

import (
	"log"
	"errors"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_service/store"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
)

type Service interface {
	GetCampaigns() ([]*campaigns.Campaign, error)
	CreateCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	EditCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error
}

type basicService struct {
	store store.Store
	iBalance balance_interface.BalanceInterface
}

func NewService(logger log.Logger, iBalance balance_interface.BalanceInterface) Service {
	var service Service
	service = basicService{store.New(logger), iBalance}
	return service
}

func (s basicService) GetCampaigns() ([]*campaigns.Campaign, error) {
	return s.store.GetCampaigns()
}

func (s basicService) EditCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	
	balance, err := s.iBalance.GetBalanceByUserID(userID)
	switch {
	case err != nil:
		return nil, err
	case balance.Amount < int64(campaign.VisitsGoal*2):
		return nil, errors.New("insufficient balance")
	}

	return s.store.EditCampaign(campaign)
}

func (s basicService) CreateCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	
	balance, err := s.iBalance.GetBalanceByUserID(userID)
	switch {
	case err != nil:
		return nil, err
	case balance.Amount < int64(campaign.VisitsGoal*2):
		return nil, errors.New("insufficient balance")
	}

	return s.store.CreateCampaign(campaign)

}

func (s basicService) PauseCampaign(id int32) error {
	return s.store.PauseCampaign(id)
}