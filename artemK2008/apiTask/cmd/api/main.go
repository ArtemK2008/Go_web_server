package main

import (
	"github.com/artemK2008/apiTask/internals/app/api"
	"log"
)

var (
	configPath string
)

func main() {
	log.Println("works")
	config := api.NewConfig()
	server := api.New(config)
	log.Fatal(server.Start())

}
