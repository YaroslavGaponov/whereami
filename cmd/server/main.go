package main

import (
	"fmt"
	"log"
	"os"

	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/internal/server"
	"github.com/YaroslavGaponov/whereami/internal/whereami"
)

var (
	fileName      string
	serverAddress string
)

func init() {
	fileName = os.Getenv("DATA_FILE")
	serverAddress = os.Getenv("SERVER_ADDRESS")
}

func main() {

	fmt.Println("whereami service")

	store := geodata.New(fileName)
	if err := store.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer store.Close()

	w := whereami.New(store)

	fmt.Print("initializing...")
	w.Initialize()
	fmt.Println("done")

	server := server.New(serverAddress, w)
	log.Fatal(server.Run())
}
