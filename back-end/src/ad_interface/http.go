package ad_interface

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
	adCodec "github.com/ruybrito106/ads-manager-services/back-end/src/ads/json"
)

const contentType = "application/json"

type adHttpClient struct {
	addr   string
	client *http.Client
}

func NewAdInterface(addr string) AdInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c AdInterface
	c = &adHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (c *adHttpClient) GetAds() ([]*ads.Ad, error) {

	getAdsAddr := c.addr + "/ads"

	res, err := c.client.Get(getAdsAddr)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from ad service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return adCodec.AdsFromJSON(body)
}

func (c *adHttpClient) CreateAd(ad *ads.Ad) (*ads.Ad, error) {

	createAdAddr := c.addr + "/ads/create"

	jsonAd, err := adCodec.AdToJSON(ad)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonAd)
	res, err := c.client.Post(createAdAddr, contentType, reader)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from ad service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return adCodec.AdFromJSON(body)

}

func (c *adHttpClient) DeleteAd(id string) error {

	deleteAdAddr := c.addr + "/ads/delete?id=" + id

	res, err := c.client.Get(deleteAdAddr)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected response from ad service")
	}

	return nil

}
