package auth_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

type AuthInterface interface {
	LoginUser(*users.User) (*users.User, error)
	GetUserByID(string) (*users.User, error)
}
