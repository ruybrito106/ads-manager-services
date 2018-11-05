package auth_controller_service

import (
	"log"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

type Service interface {
	LoginUser(*users.User) (*users.User, error)
	RegisterUser(*users.User) (*users.User, error)
}

type basicService struct {
	iAuth    auth_interface.AuthInterface
	iBalance balance_interface.BalanceInterface
}

func NewService(logger log.Logger, iAuth auth_interface.AuthInterface, iBalance balance_interface.BalanceInterface) Service {
	var service Service
	service = basicService{
		iAuth,
		iBalance,
	}
	return service
}

func (a basicService) LoginUser(user *users.User) (*users.User, error) {
	return a.iAuth.LoginUser(user)
}

func (a basicService) RegisterUser(user *users.User) (*users.User, error) {
	a.iBalance.InitBalance(user.ID)
	return a.iAuth.RegisterUser(user)
}
