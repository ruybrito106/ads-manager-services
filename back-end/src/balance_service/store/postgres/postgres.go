package postgres

import (
	"log"
	"os"

	"github.com/go-pg/pg"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balances"
)

var connection *pg.DB

type BasicDatabase interface {
	GetConnection() *pg.DB
	CloseConnection()
}

type BalanceDatabase interface {
	GetBalanceByUserID(string) (*balances.Balance, error)
	InitBalance(string) (*balances.Balance, error)
}

type basicDatabase struct {
	logger log.Logger
}

type balanceDatabase struct {
	BasicDatabase
	logger log.Logger
}

func NewDatabase(logger log.Logger) BalanceDatabase {
	var database BalanceDatabase
	basicDatabase := basicDatabase{logger}

	database = balanceDatabase{
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
		addr := os.Getenv("BALANCE_DATABASE_ADDRESS")
		port := os.Getenv("BALANCE_DATABASE_PORT")
		user := os.Getenv("BALANCE_DATABASE_USER")
		pass := os.Getenv("BALANCE_DATABASE_PASS")
		name := os.Getenv("BALANCE_DATABASE_NAME")

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

func (c balanceDatabase) GetBalanceByUserID(userID string) (*balances.Balance, error) {
	db := c.GetConnection()

	balance := balances.Balance{}

	if err := db.Model(&balance).Where("user_id = ?", userID).Select(); err != nil {
		return nil, err
	}

	return &balance, nil
}

func (c balanceDatabase) InitBalance(userID string) (*balances.Balance, error) {
	db := c.GetConnection()

	balance := balances.Balance{
		UserID: userID,
		Amount: 0,
	}

	if err := db.Insert(&balance); err != nil {
		return nil, err
	}

	return &balance, nil
}
