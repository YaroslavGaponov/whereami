package main

import (
	"fmt"
	"os"

	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/internal/whereami"
)

func main() {

	fmt.Println("whereami cli tool")

	fileName := os.Getenv("DATAFILE")
	if len(fileName) == 0 {
		fmt.Println("geodata file is not found")
	}

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

	for {
		var lat, lng float64
		fmt.Print("\nlat lng: ")
		fmt.Scan(&lat, &lng)
		city, result := w.Search(lat, lng)

		fmt.Printf("Object %s\n", result.Object.Id)
		fmt.Printf("Lat %f\n", city.Lat)
		fmt.Printf("Lng %f\n", city.Lng)
		fmt.Printf("Distance %.2f km\n", result.Distance)
		fmt.Printf("Took %v\n", result.Took)
		fmt.Printf("City %s\n", city.City)
		fmt.Printf("Country %s\n", city.Country)
	}

}
