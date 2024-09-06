package whereami

import (
	"fmt"

	"github.com/YaroslavGaponov/geosearch"
	"github.com/YaroslavGaponov/whereami/internal/geodata"
)

type WhereAmI struct {
	store  geodata.GeoData
	cities map[string]*geodata.GeoPoint
	search geosearch.GeoSearch
}

func New(store geodata.GeoData) WhereAmI {
	return WhereAmI{
		store:  store,
		cities: make(map[string]*geodata.GeoPoint),
		search: geosearch.New(5, 500),
	}
}

func (w *WhereAmI) Load() {
	fmt.Println("loading...")
	for {
		point, err := w.store.Read()
		if err != nil {
			break
		}
		w.cities[point.Id] = point
		w.search.AddObject(&geosearch.Object{Id: point.Id, Point: geosearch.Point{Latitude: point.Lat, Longitude: point.Lng}})
	}
}

func (w *WhereAmI) Search(lat, lng float64) (*geodata.GeoPoint, geosearch.Result) {
	fmt.Println("searching...")
	result := w.search.Search(geosearch.Point{Latitude: lat, Longitude: lng})
	return w.cities[result.Object.Id], result
}
