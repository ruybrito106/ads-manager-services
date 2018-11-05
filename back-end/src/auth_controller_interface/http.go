package auth_controller_interface

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/users"
	userCodec "github.com/ruybrito106/ads-manager-services/back-end/src/users/json"
)

const contentType = "application/json"

type authControllerHttpClient struct {
	addr   string
	client *http.Client
}

func NewAuthControllerInterface(addr string) AuthControllerInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c AuthControllerInterface
	c = &authControllerHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (c *authControllerHttpClient) LoginUser(user *users.User) (*users.User, error) {

	loginUserAddr := c.addr + "/controller/users/login"

	jsonUser, err := userCodec.UserToJSON(user)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonUser)
	res, err := c.client.Post(loginUserAddr, contentType, reader)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from auth service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return userCodec.UserFromJSON(body)

}

func (c *authControllerHttpClient) RegisterUser(user *users.User) (*users.User, error) {

	registerUserAddr := c.addr + "/controller/users/register"

	jsonUser, err := userCodec.UserToJSON(user)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonUser)
	res, err := c.client.Post(registerUserAddr, contentType, reader)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from auth service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return userCodec.UserFromJSON(body)

}
