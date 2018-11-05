package campaign_controller_service

import (
	"errors"
	"fmt"
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

type Service interface {
	GetCampaigns() ([]*campaigns.Campaign, error)
	CreateCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	EditCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error
}

type basicService struct {
	iCampaign campaign_interface.CampaignInterface
	iBalance  balance_interface.BalanceInterface
}

func NewService(logger log.Logger, iCampaign campaign_interface.CampaignInterface, iBalance balance_interface.BalanceInterface) Service {
	var service Service
	service = basicService{
		iCampaign,
		iBalance,
	}
	return service
}

func (s basicService) GetCampaigns() ([]*campaigns.Campaign, error) {
	return s.iCampaign.GetCampaigns()
}

func (s basicService) EditCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {

	balance, err := s.iBalance.GetBalanceByUserID(userID)
	if err != nil {
		return nil, err
	}

	if balance.Amount >= int64(campaign.VisitsGoal*2) {
		return s.iCampaign.EditCampaign(campaign)
	}

	return nil, errors.New("insufficient balance")
}

func (s basicService) CreateCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {

	balance, err := s.iBalance.GetBalanceByUserID(userID)
	fmt.Println(balance)
	if err != nil {
		return nil, err
	}

	if balance.Amount >= int64(campaign.VisitsGoal*2) {
		return s.iCampaign.CreateCampaign(campaign)
	}

	return nil, errors.New("insufficient balance")

}

func (s basicService) PauseCampaign(id int32) error {
	return s.iCampaign.PauseCampaign(id)
}
