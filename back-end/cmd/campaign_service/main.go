package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Campaign: [ack]")
	}
}

func main() {

	go alive()

	iBalance := balance_interface.NewBalanceInterface(os.Getenv("BALANCE_SERVICE_ADDRESS"))

	campaign_service.
		NewCampaignServer(
			":8081",
			log.Logger{},
			iBalance,
		).
		ListenAndServe()

}
