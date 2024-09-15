package whereami

import (
	"context"
	"errors"
	"time"

	"github.com/YaroslavGaponov/geosearch"
	"github.com/YaroslavGaponov/whereami/internal/geodata"
	"github.com/YaroslavGaponov/whereami/pkg/logger"
)

type WhereAmI struct {
	ctx         context.Context
	store       geodata.GeoData
	cities      map[string]*geodata.GeoPoint
	search      geosearch.GeoSearchFast
	initialized bool
}

type WhereAmIResponse struct {
	Id       string        `json:"id"`
	Lat      float64       `json:"lat"`
	Lng      float64       `json:"lng"`
	Distance float64       `json:"distance"`
	Took     time.Duration `json:"took"`
	City     string        `json:"city"`
	Country  string        `json:"country"`
}

var (
	notInitializedErr = errors.New("whereami is not initialized")
)

func New(ctx context.Context, store geodata.GeoData) *WhereAmI {
	return &WhereAmI{
		ctx:         ctx,
		store:       store,
		cities:      make(map[string]*geodata.GeoPoint),
		search:      geosearch.GeoSearchFastNew(500),
		initialized: false,
	}
}

func (w *WhereAmI) Initialize() {
	log := logger.GetLogger(w.ctx)
	log.Info("initializing...")
	points := 0
	for {
		point, err := w.store.Read()
		if err != nil {
			break
		}
		w.cities[point.Id] = point
		w.search.AddObject(&geosearch.Object{Id: point.Id, Point: geosearch.Point{Latitude: point.Lat, Longitude: point.Lng}})
		points++
	}
	w.initialized = true
	log.Info("%d points loaded", points)
	log.Info("done")
}

func (w *WhereAmI) IsInitialized() bool {
	return w.initialized
}

func (w *WhereAmI) Search(lat, lng float64) (*WhereAmIResponse, error) {
	if !w.initialized {
		return nil, notInitializedErr
	}
	logger.GetLogger(w.ctx).Debug("search lat=%f lng=%f", lat, lng)
	result := w.search.Search(geosearch.Point{Latitude: lat, Longitude: lng})
	city := w.cities[result.Object.Id]
	return &WhereAmIResponse{
		Id:       result.Object.Id,
		Lat:      city.Lat,
		Lng:      city.Lng,
		Distance: result.Distance,
		Took:     result.Took,
		City:     city.City,
		Country:  city.Country,
	}, nil
}
