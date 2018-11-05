package balance_interface

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balances"
	balanceCodec "github.com/ruybrito106/ads-manager-services/back-end/src/balances/json"
)

const contentType = "application/json"

type balanceHttpClient struct {
	addr   string
	client *http.Client
}

func NewBalanceInterface(addr string) BalanceInterface {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	var c BalanceInterface
	c = &balanceHttpClient{
		addr,
		&http.Client{
			Transport: tr,
			Timeout:   time.Minute,
		},
	}

	return c
}

func (c *balanceHttpClient) GetBalanceByUserID(userID string) (*balances.Balance, error) {

	getBalanceByUserIDAddr := c.addr + "/balances?id=" + userID

	res, err := c.client.Get(getBalanceByUserIDAddr)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from balance service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return balanceCodec.BalanceFromJSON(body)

}

func (c *balanceHttpClient) InitBalance(userID string) (*balances.Balance, error) {

	initBalanceByUserIDAddr := c.addr + "/balances/init?id=" + userID

	res, err := c.client.Get(initBalanceByUserIDAddr)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response from balance service")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return balanceCodec.BalanceFromJSON(body)

}
