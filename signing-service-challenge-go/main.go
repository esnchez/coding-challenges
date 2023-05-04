package main

import (
	"log"

	"github.com/esnchez/coding-challenges/signing-service-challenge/api"
	"github.com/esnchez/coding-challenges/signing-service-challenge/persistence"
	"github.com/esnchez/coding-challenges/signing-service-challenge/service"
)

const (
	ListenAddress = ":8080"
	// TODO: add further configuration parameters here ...
)

func main() {

	//init repo & svc
	store := persistence.NewMemStore()
	svc := service.NewSignatureService(store)

	server := api.NewServer(ListenAddress, svc)

	if err := server.Run(); err != nil {
		log.Fatal("Could not start server on ", ListenAddress)
	}
}
