package controller

import (
	"encoding/json"
	"log"
	"net/http"

	house "github.com/massintha/house-finder/src/house-finder/models"
	"github.com/massintha/house-finder/src/house-finder/providers"
	"github.com/massintha/house-finder/src/house-finder/services/geotag"
	"github.com/massintha/house-finder/src/house-finder/store"
)

// ServerHouseList it's in the name
func ServerHouseList(store *store.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		houseList, err := store.GetAllHouses()
		if err != nil {
			log.Fatal("could not get house list")
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			jsonHouse, _ := json.Marshal(houseList)

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonHouse)
		}
	}
}

// GetBoundingBox it's in the name
func GetBoundingBox(store *store.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		houseEnvelope, _ := store.GetBoundingBox()
		coordinates := houseEnvelope.Coordinates[0]
		firstPoint := coordinates[0]
		lastPoint := coordinates[2]

		boundingBox := make(map[string]float64)
		boundingBox["x1"] = firstPoint[0]
		boundingBox["y1"] = firstPoint[1]
		boundingBox["x2"] = lastPoint[0]
		boundingBox["y2"] = lastPoint[1]

		output := make(map[string]map[string]float64)
		output["boundingBox"] = boundingBox
		jsonHouse, _ := json.Marshal(output)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonHouse)
	}
}

// UpdateHouse it's in the name
func UpdateHouse(store *store.Store) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var house house.House

		err := json.NewDecoder(r.Body).Decode(&house)
		if err != nil {
			log.Fatal("could not marshal house to json")
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			store.UpdateHouse(house)
		}

		w.Header().Set("Content-Type", "application/json")
	}
}

// SyncHouses it's in the name
func SyncHouses(store *store.Store, geoTag *geotag.GeoTag) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		providers.SyncHouses(store, geoTag)

		w.Header().Set("Content-Type", "application/json")
	}
}
