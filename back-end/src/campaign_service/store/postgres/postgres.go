package postgres

import (
	"log"
	"os"
	"fmt"

	"github.com/go-pg/pg"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaigns"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
}

type CampaignDatabase interface {
	GetCampaigns() ([]*campaigns.Campaign, error)
	EditCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	CreateCampaign(*campaigns.Campaign) (*campaigns.Campaign, error)
	PauseCampaign(id int32) error
}

type basicDatabase struct {
	logger log.Logger
}

type campaignDatabase struct {
	BasicDatabase
	logger log.Logger
}

func NewDatabase(logger log.Logger) CampaignDatabase {
	var database CampaignDatabase
	basicDatabase := basicDatabase{logger}

	database = campaignDatabase{
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
		addr := os.Getenv("CAMPAIGN_DATABASE_ADDRESS")
		port := os.Getenv("CAMPAIGN_DATABASE_PORT")
		user := os.Getenv("CAMPAIGN_DATABASE_USER")
		pass := os.Getenv("CAMPAIGN_DATABASE_PASS")
		name := os.Getenv("CAMPAIGN_DATABASE_NAME")

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

func (c campaignDatabase) CreateCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	db := c.GetConnection()

	if err := db.Insert(campaign); err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c campaignDatabase) EditCampaign(campaign *campaigns.Campaign) (*campaigns.Campaign, error) {
	db := c.GetConnection()

	if err := db.Update(campaign); err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c campaignDatabase) GetCampaigns() ([]*campaigns.Campaign, error) {
	db := c.GetConnection()

	campaigns := []*campaigns.Campaign{}

	if err := db.Model(&campaigns).Select(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return campaigns, nil
}

func (c campaignDatabase) PauseCampaign(id int32) error {
	db := c.GetConnection()

	campaign := campaigns.Campaign{}

	if _, err := db.Model(&campaign).Set("status = ?", campaigns.StatusPaused).Where("id = ?", id).Update(); err != nil {
		return err
	}

	return nil
}
