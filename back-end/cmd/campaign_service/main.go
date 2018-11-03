package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/campaign_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Campaign: [ack]")
	}
}

func main() {

	go alive()

	campaign_service.
		NewCampaignServer(":8081", log.Logger{}).
		ListenAndServe()

}
