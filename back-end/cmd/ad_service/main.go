package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/ad_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Ad: [ack]")
	}
}

func main() {

	go alive()

	ad_service.
		NewAdServer(
			":8085",
			log.Logger{},
		).
		ListenAndServe()

}
