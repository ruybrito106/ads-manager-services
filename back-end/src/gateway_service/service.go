package gateway_service

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_controller_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

type Service interface {
	// Campaign Controller Service Methods
	GetCampaigns() ([]*campaigns.Campaign, error)
	CreateCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	EditCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error

	// Auth Service Methods
	LoginUser(*users.User) (*users.User, error)
}

type basicService struct {
	IAuth               auth_interface.AuthInterface
	ICampaignController campaign_controller_interface.CampaignControllerInterface
}

func NewService(iCampaignController campaign_controller_interface.CampaignControllerInterface, iAuth auth_interface.AuthInterface) Service {
	var service Service
	service = basicService{
		ICampaignController: iCampaignController,
		IAuth:               iAuth,
	}
	return service
}

func (s basicService) LoginUser(user *users.User) (*users.User, error) {
	return s.IAuth.LoginUser(user)
}

func (s basicService) GetCampaigns() ([]*campaigns.Campaign, error) {
	return s.ICampaignController.GetCampaigns()
}

func (s basicService) EditCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaignController.EditCampaign(userID, campaign)
}

func (s basicService) CreateCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaignController.CreateCampaign(userID, campaign)
}

func (s basicService) PauseCampaign(id int32) error {
	return s.ICampaignController.PauseCampaign(id)
}
