package store

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_service/store/postgres"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balances"
)

type Store interface {
	GetBalanceByUserID(id string) (*balances.Balance, error)
}

type basicStore struct {
	logger log.Logger
}

func New(logger log.Logger) basicStore {
	return basicStore{logger}
}

func (s basicStore) GetBalanceByUserID(userID string) (*balances.Balance, error) {
	db := postgres.NewDatabase(s.logger)
	return db.GetBalanceByUserID(userID)
}
