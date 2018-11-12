package place_interface

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
	placeCodec "github.com/ruybrito106/ads-manager-services/back-end/src/places/json"
)

const contentType = "application/json"

type placeHttpClient struct {
	addr   string
	client *http.Client
}

func NewPlaceInterface(addr string) PlaceInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c PlaceInterface
	c = &placeHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (c *placeHttpClient) GetPlaces() ([]*places.Place, error) {

	getPlacesAddr := c.addr + "/places"

	res, err := c.client.Get(getPlacesAddr)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from place service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return placeCodec.PlacesFromJSON(body)
}

func (c *placeHttpClient) CreatePlace(place *places.Place) (*places.Place, error) {

	createPlaceAddr := c.addr + "/places/create"

	jsonPlace, err := placeCodec.PlaceToJSON(place)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonPlace)
	res, err := c.client.Post(createPlaceAddr, contentType, reader)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from place service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return placeCodec.PlaceFromJSON(body)

}

func (c *placeHttpClient) DeletePlace(id string) error {

	deletePlaceAddr := c.addr + "/places/delete?id=" + id

	res, err := c.client.Get(deletePlaceAddr)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected response from place service")
	}

	return nil

}
