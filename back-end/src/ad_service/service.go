package ad_service

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ad_service/store"
	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
)

type Service interface {
	CreateAd(*ads.Ad) (*ads.Ad, error)
	DeleteAd(string) error
	GetAds(string) ([]*ads.Ad, error)
}

type basicService struct {
	store store.Store
}

func NewService(logger log.Logger) Service {
	var service Service
	service = basicService{store.New(logger)}
	return service
}

func (s basicService) CreateAd(ad *ads.Ad) (*ads.Ad, error) {
	return s.store.CreateAd(ad)
}

func (s basicService) DeleteAd(id string) error {
	return s.store.DeleteAd(id)
}

func (s basicService) GetAds(userID string) ([]*ads.Ad, error) {
	return s.store.GetAds(userID)
}
