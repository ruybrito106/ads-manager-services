package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ruybrito106/ads-manager-services/back-end/src/balance_service"
)

func alive() {
	for range time.Tick(30 * time.Second) {
		fmt.Println("Balance: [ack]")
	}
}

func main() {

	go alive()

	balance_service.
		NewBalanceServer(":8082", log.Logger{}).
		ListenAndServe()

}
