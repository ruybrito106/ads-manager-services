package balance_service

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_service/store"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balances"
)

type Service interface {
	GetBalanceByUserID(string) (*balances.Balance, error)
	InitBalance(string) (*balances.Balance, error)
}

type basicService struct {
	store store.Store
}

func NewService(logger log.Logger) Service {
	var service Service
	service = basicService{store.New(logger)}
	return service
}

func (s basicService) GetBalanceByUserID(userID string) (*balances.Balance, error) {
	return s.store.GetBalanceByUserID(userID)
}

func (s basicService) InitBalance(userID string) (*balances.Balance, error) {
	return s.store.InitBalance(userID)
}
