package campaign_controller_interface

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

type campaignControllerHttpClient struct {
	addr   string
	client *http.Client
}

func NewCampaignControllerInterface(addr string) CampaignControllerInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c CampaignControllerInterface
	c = &campaignControllerHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (c *campaignControllerHttpClient) GetCampaigns() ([]*campaigns.Campaign, error) {

	getCampaignsAddr := c.addr + "/controller/campaigns"

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

func (c *campaignControllerHttpClient) CreateCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {

	createCampaignAddr := c.addr + "/controller/campaigns/create?id=" + userID

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

func (c *campaignControllerHttpClient) EditCampaign(userID string, campaign *campaigns.Campaign) (*campaigns.Campaign, error) {

	editCampaignAddr := c.addr + "/controller/campaigns/edit?id=" + userID

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

func (c *campaignControllerHttpClient) PauseCampaign(id int32) error {

	pauseCampaignAddr := strings.Join(
		[]string{
			c.addr,
			"/controller/campaigns/pause?id=",
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
