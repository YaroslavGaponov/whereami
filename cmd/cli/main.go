package main

import (
	"fmt"
	"os"

	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/internal/whereami"
)

var (
	fileName string
)

func init() {
	fileName = os.Getenv("DATA_FILE")
}

func main() {

	fmt.Println("whereami cli tool")

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
		result := w.Search(lat, lng)

		fmt.Printf("Object %s\n", result.Id)
		fmt.Printf("Lat %f\n", result.Lat)
		fmt.Printf("Lng %f\n", result.Lng)
		fmt.Printf("Distance %.2f km\n", result.Distance)
		fmt.Printf("Took %v\n", result.Took)
		fmt.Printf("City %s\n", result.City)
		fmt.Printf("Country %s\n", result.Country)
	}

}
