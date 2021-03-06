package campaign_interface

import (
	"testing"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
	"github.com/stretchr/testify/assert"
)

func TestCreateCampaign(t *testing.T) {

	fixture := &campaigns.Campaign{
		Name:       "xibliu",
		VisitsGoal: 100,
		Status:     campaigns.StatusActive,
	}

	c := NewCampaignInterface("http://localhost:8081")

	created, err := c.CreateCampaign("132", fixture)
	if !assert.NoError(t, err) {
		return
	}

	fixture.ID = created.ID
	assert.Equal(t, fixture, created)

}
