package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"

	"github.com/ruybrito106/ads-manager-services/back-end/src/places"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
}

type PlaceDatabase interface {
	CreatePlace(*places.Place) (*places.Place, error)
	DeletePlace(id string) error
	GetPlaces() ([]*places.Place, error)
}

type basicDatabase struct {
	logger log.Logger
}

type placeDatabase struct {
	BasicDatabase
	logger log.Logger
}

func NewDatabase(logger log.Logger) PlaceDatabase {
	var database PlaceDatabase
	basicDatabase := basicDatabase{logger}

	database = placeDatabase{
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
		addr := os.Getenv("PLACE_DATABASE_ADDRESS")
		port := os.Getenv("PLACE_DATABASE_PORT")
		user := os.Getenv("PLACE_DATABASE_USER")
		pass := os.Getenv("PLACE_DATABASE_PASS")
		name := os.Getenv("PLACE_DATABASE_NAME")

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

func (c placeDatabase) CreatePlace(place *places.Place) (*places.Place, error) {
	db := c.GetConnection()

	if err := db.Insert(place); err != nil {
		return nil, err
	}

	return place, nil
}

func (c placeDatabase) GetPlaces() ([]*places.Place, error) {
	db := c.GetConnection()

	places := []*places.Place{}

	if err := db.Model(&places).Select(); err != nil {
		return nil, err
	}

	return places, nil
}

func (c placeDatabase) DeletePlace(id string) error {
	db := c.GetConnection()

	place := places.Place{}

	if _, err := db.Model(&place).Where("id = ?", id).Delete(); err != nil {
		return err
	}

	return nil
}
