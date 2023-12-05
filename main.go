package main

import (
	"fmt"

	"github.com/tigertony2536/go-line-notify/config"
)

func main() {
	// cfg := config.GetConfig("config/config.yaml")

	config := config.GetConfig()
	fmt.Println(config.DB)
	fmt.Println(config.Token)
}
