package gateway_service

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

type Service interface {
	// Campaign Service Methods
	GetCampaigns() ([]*campaigns.Campaign, error)
	CreateCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	EditCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error

	// Auth Service Methods
	LoginUser(*users.User) (*users.User, error)
	RegisterUser(*users.User) (*users.User, error)
}

type basicService struct {
	IAuth     auth_interface.AuthInterface
	ICampaign campaign_interface.CampaignInterface
}

func NewService(iCampaign campaign_interface.CampaignInterface, iAuth auth_interface.AuthInterface) Service {
	var service Service
	service = basicService{
		ICampaign: iCampaign,
		IAuth:     iAuth,
	}
	return service
}

func (s basicService) LoginUser(user *users.User) (*users.User, error) {
	return s.IAuth.LoginUser(user)
}

func (s basicService) RegisterUser(user *users.User) (*users.User, error) {
	return s.IAuth.RegisterUser(user)
}

func (s basicService) GetCampaigns() ([]*campaigns.Campaign, error) {
	return s.ICampaign.GetCampaigns()
}

func (s basicService) EditCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaign.EditCampaign(userID, campaign)
}

func (s basicService) CreateCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	return s.ICampaign.CreateCampaign(userID, campaign)
}

func (s basicService) PauseCampaign(id int32) error {
	return s.ICampaign.PauseCampaign(id)
}
