package main

import (
	"net/http"

	_ "github.com/jackc/pgx/v4"
	"github.com/massintha/house-finder/src/house-finder/controller"
	"github.com/massintha/house-finder/src/house-finder/services/geotag"
	"github.com/massintha/house-finder/src/house-finder/store"
)

func main() {
	geoTag := geotag.BuildGeoTag()
	store, err := store.SetupStore()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/houses", controller.ServerHouseList(store))
	http.HandleFunc("/bounding-box", controller.GetBoundingBox(store))
	http.HandleFunc("/house", controller.UpdateHouse(store))
	http.HandleFunc("/sync-houses", controller.SyncHouses(store, geoTag))

	http.ListenAndServe(":8080", nil)

}
