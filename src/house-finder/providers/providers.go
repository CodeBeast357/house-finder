package providers

import (
	"fmt"
	"log"

	house "github.com/massintha/house-finder/src/house-finder/models"
	"github.com/massintha/house-finder/src/house-finder/providers/centris"
	"github.com/massintha/house-finder/src/house-finder/providers/duproprio"
	"github.com/massintha/house-finder/src/house-finder/providers/provider"
	"github.com/massintha/house-finder/src/house-finder/providers/remax"
	"github.com/massintha/house-finder/src/house-finder/services/geotag"
	"github.com/massintha/house-finder/src/house-finder/services/mapbox/geocode"
	"github.com/massintha/house-finder/src/house-finder/store"
)

// SyncHouses it's in the name
func SyncHouses(store *store.Store, geoTag *geotag.GeoTag) {
	houses := make([]*house.House, 0)
	fmt.Println("Starting sync")

	// arrondissementFilters := []provider.ArrondissementFilter{provider.Ahunstic, provider.MontRoyal, provider.Rosemont, provider.Villeray, provider.TroisRivieres}
	arrondissementFilters := []provider.ArrondissementFilter{provider.Ahunstic, provider.MontRoyal, provider.Rosemont, provider.Villeray}
	for _, arrondissement := range arrondissementFilters {
		fmt.Println("Sync arrondissement ", arrondissement)
		fmt.Println("Sync centris")
		houses = append(houses, centris.GetHouses(arrondissement)...)
		fmt.Println("Sync duproprio")
		houses = append(houses, duproprio.GetHouses(arrondissement)...)
		fmt.Println("Sync remax")
		houses = append(houses, remax.GetHouses(arrondissement)...)
	}

	fmt.Println("Fetching addresses coordinates")
	for i := range houses {
		house := houses[i]
		forwardOpts := geocode.ForwardRequestOpts{
			BBox:  provider.ArrondissementBBox[house.Arrondissement],
			Limit: 3,
		}
		forward, err := geocode.Forward(house, &forwardOpts)
		if err != nil {
			log.Fatalln(err)
		}

		position := forward.Features[0].Center
		for _, point := range forward.Features {
			center := point.Center
			if geoTag.IsPointInArrondissement(house.Arrondissement, center) {
				position = center
				break
			}
		}
		house.SetCoordinates(position)
		house.SetIsInSweetSpot(geoTag.IsInSweetSpot(position))
	}

	fmt.Println("Saving to DB")
	err := store.SaveHouseList(houses)
	if err != nil {
		log.Fatalln("Error while saving to DB", err)
	}
	fmt.Println("Sync done")
}
