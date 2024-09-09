package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/internal/server"
	"github.com/YaroslavGaponov/whereami/internal/whereami"
)

var (
	fileName string
	port     int
)

func init() {
	fileName = os.Getenv("DATAFILE")
	if len(fileName) == 0 {
		log.Fatal("geodata file is not found")
	}

	var err error
	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
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

	server := server.New(port, w)
	server.Run()
}
