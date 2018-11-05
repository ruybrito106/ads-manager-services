package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_interface"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_controller_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Campaign Controller: [ack]")
	}
}

func main() {

	go alive()

	iCampaign := campaign_interface.NewCampaignInterface(os.Getenv("CAMPAIGN_SERVICE_ADDRESS"))
	iBalance := balance_interface.NewBalanceInterface(os.Getenv("BALANCE_SERVICE_ADDRESS"))

	campaign_controller_service.
		NewCampaignControllerServer(
			":8083",
			log.Logger{},
			iCampaign,
			iBalance,
		).
		ListenAndServe()

}
