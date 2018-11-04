package campaign_interface

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
	campaignCodec "github.com/ruybrito106/ads-manager-services/back-end/src/campaigns/json"
)

const contentType = "application/json"

type campaignHttpClient struct {
	addr   string
	client *http.Client
}

func NewCampaignInterface(addr string) CampaignInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c CampaignInterface
	c = &campaignHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (c *campaignHttpClient) GetCampaigns() ([]*campaigns.Campaign, error) {

	getCampaignsAddr := c.addr + "/campaigns"

	res, err := c.client.Get(getCampaignsAddr)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from campaign service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return campaignCodec.CampaignsFromJSON(body)
}

func (c *campaignHttpClient) CreateCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {

	createCampaignAddr := c.addr + "/campaigns/create"

	jsonCampaign, err := campaignCodec.CampaignToJSON(campaign)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonCampaign)
	res, err := c.client.Post(createCampaignAddr, contentType, reader)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from campaign service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return campaignCodec.CampaignFromJSON(body)

}

func (c *campaignHttpClient) EditCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {

	editCampaignAddr := c.addr + "/campaigns/edit"

	jsonCampaign, err := campaignCodec.CampaignToJSON(campaign)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonCampaign)
	res, err := c.client.Post(editCampaignAddr, contentType, reader)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from campaign service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return campaignCodec.CampaignFromJSON(body)

}


func (c *campaignHttpClient) PauseCampaign(id int32) error {

	pauseCampaignAddr := strings.Join(
		[]string{
			c.addr,
			"/campaigns/pause?id=",
			strconv.Itoa(int(id)),
		},
		"",
	)

	res, err := c.client.Get(pauseCampaignAddr)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected response from campaign service")
	}

	return nil

}
