package main

import (
	"fmt"
	"log"

	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

func main() {
	cfg := config.GetConfig()
	db := model.GetDB(cfg.DB)

	noti, err := db.GetByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(noti)
}
