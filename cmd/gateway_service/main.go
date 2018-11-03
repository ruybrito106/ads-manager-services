package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ruybrito106/ads-manager-services/src/campaign_interface"
	"github.com/ruybrito106/ads-manager-services/src/gateway_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Gateway: [ack]")
	}
}

func main() {

	go alive()

	iCampaign := campaign_interface.NewCampaignInterface(os.Getenv("CAMPAIGN_SERVICE_ADDRESS"))

	gateway_service.
		NewGatewayServer(":8080", iCampaign).
		ListenAndServe()

}
