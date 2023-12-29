package main

import (
	"fmt"
	"log"

	"github.com/tigertony2536/go-line-notify/controller"
)

func main() {
	// server := controller.NewServer(":8080")
	// log.Fatal(server.Start())

	weekNoti, err := controller.GetWeeklyNoti()
	if err != nil {
		log.Fatal(err)
	}
	dailyNoti, err := controller.GetDailyNoti()
	if err != nil {
		log.Fatal(err)
	}
	if weekNoti.Notifications != nil {
		resp, err := controller.SendNotification(weekNoti)
		fmt.Print(resp)
		if err != nil {
			log.Fatal(err)
		}
	}

	if dailyNoti.Notifications != nil {
		resp, err := controller.SendNotification(weekNoti)
		fmt.Print(resp)
		if err != nil {
			log.Fatal(err)
		}
	}
}
