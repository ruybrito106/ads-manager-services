package balance_interface

import (
	"github.com/ruybrito106/ads-manager-services/back-end/src/balances"
)

type BalanceInterface interface {
	GetBalanceByUserID(string) (*balances.Balance, error)
}
