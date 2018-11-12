package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/place_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Place: [ack]")
	}
}

func main() {

	go alive()

	place_service.
		NewPlaceServer(
			":8084",
			log.Logger{},
		).
		ListenAndServe()

}
