package json

import (
	"encoding/json"

	"github.com/ruybrito106/ads-manager-services/src/campaigns"
)

func CampaignToJSON(campaign *campaigns.Campaign) ([]byte, error) {
	return json.Marshal(&campaign)
}

func CampaignFromJSON(encoded []byte) (*campaigns.Campaign, error) {
	campaign := campaigns.Campaign{}
	if err := json.Unmarshal(encoded, &campaign); err != nil {
		return nil, err
	}
	return &campaign, nil
}
