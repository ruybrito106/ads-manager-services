package campaign_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

type CampaignInterface interface {
	GetCampaigns() ([]*campaigns.Campaign, error)
	EditCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	CreateCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(int32) error
}