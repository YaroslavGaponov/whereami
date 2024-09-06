package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/internal/whereami"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("help: whereami {latitude} {longitude}")
		fmt.Println("example: whereami 53.876592, 14.266755")
		os.Exit(1)
	}
	lat, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("latitude is incorrect")
		os.Exit(1)
	}
	lng, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("longitude is incorrect")
		os.Exit(1)
	}

	fileName, err := getGeoDateFileName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	store := geodata.New(fileName)
	if err := store.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer store.Close()

	w := whereami.New(store)

	w.Load()

	city, result := w.Search(lat,lng)

	fmt.Println("Result:")
	fmt.Printf("\tObject %s\n", result.Object.Id)
	fmt.Printf("\tDistance %.2f km\n", result.Distance)
	fmt.Printf("\tTook %v\n", result.Took)
	fmt.Printf("\tCity %s\n", city.City)
	fmt.Printf("\tCountry %s\n", city.Country)

}

func getGeoDateFileName() (string, error) {
	_, fileName, _, _ := runtime.Caller(0)
	geoDataFolder, err := filepath.Abs(filepath.Dir(fileName) + "../../../geodata")
	if err != nil {
		return "", err
	}
	return geoDataFolder + "/worldcities.zip@worldcities.csv", nil
}
