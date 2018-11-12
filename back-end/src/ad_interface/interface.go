package ad_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
)

type AdInterface interface {
	CreateAd(*ads.Ad) (*ads.Ad, error)
	DeleteAd(string) error
	GetAds() ([]*ads.Ad, error)
}
