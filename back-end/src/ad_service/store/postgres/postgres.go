package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ads"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
}

type AdDatabase interface {
	CreateAd(*ads.Ad) (*ads.Ad, error)
	DeleteAd(id string) error
	GetAds(string) ([]*ads.Ad, error)
}

type basicDatabase struct {
	logger log.Logger
}

type adDatabase struct {
	BasicDatabase
	logger log.Logger
}

func NewDatabase(logger log.Logger) AdDatabase {
	var database AdDatabase
	basicDatabase := basicDatabase{logger}

	database = adDatabase{
		basicDatabase,
		logger,
	}

	return database
}

func (d basicDatabase) CloseConnection() {
	if connection != nil {
		d.logger.Println("message", "Closing Postgres session")
		connection.Close()
	}
}

func (d basicDatabase) GetConnection() *pg.DB {
	if connection == nil {
		addr := os.Getenv("AD_DATABASE_ADDRESS")
		port := os.Getenv("AD_DATABASE_PORT")
		user := os.Getenv("AD_DATABASE_USER")
		pass := os.Getenv("AD_DATABASE_PASS")
		name := os.Getenv("AD_DATABASE_NAME")

		connection = pg.Connect(&pg.Options{
			User:     user,
			Password: pass,
			Database: name,
			Addr:     addr + ":" + port,
			PoolSize: 30,
		})
	}

	return connection
}

func (c adDatabase) CreateAd(ad *ads.Ad) (*ads.Ad, error) {
	db := c.GetConnection()

	if err := db.Insert(ad); err != nil {
		return nil, err
	}

	return ad, nil
}

func (c adDatabase) GetAds(userID string) ([]*ads.Ad, error) {
	db := c.GetConnection()

	ads := []*ads.Ad{}

	if err := db.Model(&ads).Where("user_id = ?", userID).Select(); err != nil {
		return nil, err
	}

	return ads, nil
}

func (c adDatabase) DeleteAd(id string) error {
	db := c.GetConnection()

	ad := ads.Ad{}

	if _, err := db.Model(&ad).Where("id = ?", id).Delete(); err != nil {
		return err
	}

	return nil
}
