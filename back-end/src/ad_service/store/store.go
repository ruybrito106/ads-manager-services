package store

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ad_service/store/postgres"
	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
)

type Store interface {
	CreateAd(*ads.Ad) (*ads.Ad, error)
	DeleteAd(string) error
	GetAds(string) ([]*ads.Ad, error)
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) basicStore {
	return basicStore{logger}
}

func (s basicStore) CreateAd(ad *ads.Ad) (*ads.Ad, error) {
	db := postgres.NewDatabase(s.logger)
	return db.CreateAd(ad)
}

func (s basicStore) DeleteAd(id string) error {
	db := postgres.NewDatabase(s.logger)
	return db.DeleteAd(id)
}

func (s basicStore) GetAds(userID string) ([]*ads.Ad, error) {
	db := postgres.NewDatabase(s.logger)
	return db.GetAds(userID)
}
