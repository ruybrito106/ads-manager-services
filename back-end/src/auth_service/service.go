package auth_service

import (
	"errors"
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

type Service interface {
	LoginUser(*users.User) (*users.User, error)
	RegisterUser(*users.User) (*users.User, error)
}

type basicService struct {
	iBalance balance_interface.BalanceInterface
}

func NewService(logger log.Logger, iBalance balance_interface.BalanceInterface) Service {
	var service Service
	service = basicService{iBalance}
	return service
}

func (s basicService) LoginUser(user *users.User) (*users.User, error) {

	// Should perform http request to external auth subsystem
	if user.IsSuperuser() {
		return user, nil
	}

	return nil, errors.New("login failed")

}

func (s basicService) RegisterUser(user *users.User) (*users.User, error) {

	// Should perform http request to external auth subsystem
	s.iBalance.InitBalance(user.ID)
	return user, nil

}
