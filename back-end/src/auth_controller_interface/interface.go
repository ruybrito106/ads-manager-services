package auth_controller_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

type AuthControllerInterface interface {
	LoginUser(*users.User) (*users.User, error)
	RegisterUser(*users.User) (*users.User, error)
}
