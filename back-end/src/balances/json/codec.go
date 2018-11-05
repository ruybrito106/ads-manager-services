package json

import (
	"encoding/json"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balances"
)

func BalanceToJSON(balance *balances.Balance) ([]byte, error) {
	return json.Marshal(&balance)
}

func BalanceFromJSON(encoded []byte) (*balances.Balance, error) {
	balance := balances.Balance{}
	if err := json.Unmarshal(encoded, &balance); err != nil {
		return nil, err
	}
	return &balance, nil
}
