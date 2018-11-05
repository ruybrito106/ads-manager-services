package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_controller_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Auth Controller: [ack]")
	}
}

func main() {

	go alive()

	iAuth := auth_interface.NewAuthInterface(os.Getenv("AUTH_SERVICE_ADDRESS"))
	iBalance := balance_interface.NewBalanceInterface(os.Getenv("BALANCE_SERVICE_ADDRESS"))

	auth_controller_service.
		NewAuthControllerServer(
			":8084",
			log.Logger{},
			iAuth,
			iBalance,
		).
		ListenAndServe()

}
