package auth_interface

import (
	"errors"
	"net/http"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
)

const contentType = "application/json"

type authHttpClient struct {
	addr   string
	client *http.Client
}

func NewAuthInterface(addr string) AuthInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c AuthInterface
	c = &authHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (a *authHttpClient) LoginUser(user *users.User) (*users.User, error) {

	// Should perform http request to external auth subsystem
	if user.IsSuperuser() {
		return user, nil
	}

	return nil, errors.New("login failed")

}

func (a *authHttpClient) RegisterUser(user *users.User) (*users.User, error) {

	// Should perform http request to external auth subsystem
	return user, nil

}
