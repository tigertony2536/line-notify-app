package main

import (
	"fmt"
	"log"

	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/controller"
	"github.com/tigertony2536/go-line-notify/model"
)

func main() {
	// server := controller.NewServer(":8080")
	// log.Fatal(server.Start())
	config := config.GetConfig()
	db := model.GetDB(config.DB)
	noti, err := db.GetByDate("2023-12-01", "2023-12-30")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := controller.SendNotification(noti[0])
	fmt.Print(resp)
	log.Fatal(err)
}
