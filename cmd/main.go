package main

import (
	"fmt"
	"time"

	"github.com/gau31rav/sentryintegration/logging"
	"github.com/gau31rav/sentryintegration/pkg/service/order"
)

func main() {
	logger := logging.NewAppLogger()
	orderSVC := order.NewService(logger)
	err := orderSVC.PlaceOrder()
	if err != nil {
		fmt.Println("printing before logging")
	}
	time.Sleep(10 * time.Second)
}
