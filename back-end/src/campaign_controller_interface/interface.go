package campaign_controller_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

type CampaignControllerInterface interface {
	GetCampaigns() ([]*campaigns.Campaign, error)
	EditCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	CreateCampaign(string, *campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error
}
