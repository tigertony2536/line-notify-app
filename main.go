package main

import (
	"log"

	"github.com/tigertony2536/go-line-notify/controller"
)

func main() {
	server := controller.NewServer(":8080")
	log.Fatal(server.Start())
}
