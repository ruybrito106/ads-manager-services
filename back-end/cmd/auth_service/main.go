package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_service"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Auth: [ack]")
	}
}

func main() {

	go alive()

	iBalance := balance_interface.NewBalanceInterface(os.Getenv("BALANCE_SERVICE_ADDRESS"))

	auth_service.
		NewAuthServer(
			":8083",
			log.Logger{},
			iBalance,
		).
		ListenAndServe()

}
