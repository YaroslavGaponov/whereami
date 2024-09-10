package main

import (
	"context"
	"fmt"
	"os"

	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/internal/server"
	"github.com/YaroslavGaponov/whereami/internal/whereami"
	"github.com/YaroslavGaponov/whereami/pkg/logger"
)

var (
	logLevel      string
	fileName      string
	serverAddress string
)

func init() {
	logLevel = os.Getenv("LOG_LEVEL")
	fileName = os.Getenv("DATA_FILE")
	serverAddress = os.Getenv("SERVER_ADDRESS")
}

func main() {
	log := logger.New()
	log.SetLogLevel(logLevel)

	log.Info("whereami service")

	store := geodata.New(fileName)
	if err := store.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer store.Close()

	ctx := context.WithValue(context.Background(),"logger", log)

	w := whereami.New(ctx,store)

	log.Info("initializing...")
	w.Initialize()
	
	log.Info("searvice is ready...")
	server := server.New(serverAddress, w)
	if err := server.Run(); err != nil {
		log.Fatal("%v",err)
	}
}
