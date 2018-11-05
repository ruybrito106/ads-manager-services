package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/auth_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_controller_interface"
	"github.com/ruybrito106/ads-manager-services/back-end/src/gateway_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Gateway: [ack]")
	}
}

func main() {

	go alive()

	iAuth := auth_interface.NewAuthInterface(os.Getenv("AUTH_SERVICE_ADDRESS"))
	iCampaignController := campaign_controller_interface.NewCampaignControllerInterface(os.Getenv("CAMPAIGN_CONTROLLER_SERVICE_ADDRESS"))

	gateway_service.
		NewGatewayServer(
			":8080",
			iCampaignController,
			iAuth,
		).
		ListenAndServe()

}
