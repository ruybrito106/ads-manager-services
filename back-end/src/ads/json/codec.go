package json

import (
	"encoding/json"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
)

func AdToJSON(ad *ads.Ad) ([]byte, error) {
	return json.Marshal(&ad)
}

func AdFromJSON(encoded []byte) (*ads.Ad, error) {
	ad := ads.Ad{}
	if err := json.Unmarshal(encoded, &ad); err != nil {
		return nil, err
	}
	return &ad, nil
}

func AdsToJSON(ads []*ads.Ad) ([]byte, error) {
	return json.Marshal(&ads)
}

func AdsFromJSON(encoded []byte) ([]*ads.Ad, error) {
	ads := []*ads.Ad{}
	if err := json.Unmarshal(encoded, &ads); err != nil {
		return nil, err
	}
	return ads, nil
}
