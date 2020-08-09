package geotag

import (
	"fmt"
	"io/ioutil"
	"log"

	geo "github.com/kellydunn/golang-geo"
	"github.com/massintha/house-finder/src/house-finder/providers/provider"
	geojson "github.com/paulmach/go.geojson"
)

// GeoTag GeoTag
type GeoTag struct {
	sweetSpot *geo.Polygon
	ahuntsic  *geo.Polygon
	mauricie  *geo.Polygon
	montRoyal *geo.Polygon
	rosemont  *geo.Polygon
	villeray  *geo.Polygon
}

func readFile(fileName string) []byte {
	content, err := ioutil.ReadFile(fmt.Sprintf("/house-finder/data/%s.json", fileName))
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func buildPolygon(fileName string) *geo.Polygon {
	polygonStr := readFile(fileName)
	fc, err := geojson.UnmarshalFeatureCollection(polygonStr)
	if err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}

	emptyPoints := make([]*geo.Point, 0)
	polygon := geo.NewPolygon(emptyPoints)
	for _, coords := range fc.Features[0].Geometry.Polygon[0] {
		latitude := coords[0]
		longitude := coords[1]

		point := geo.NewPoint(latitude, longitude)
		polygon.Add(point)
	}

	return polygon
}

// IsInSweetSpot IsInSweetSpot
func (geoTag *GeoTag) IsInSweetSpot(coordinates []float64) bool {
	point := geo.NewPoint(coordinates[0], coordinates[1])
	return geoTag.sweetSpot.Contains(point)
}

// IsPointInArrondissement IsPointInArrondissement
func (geoTag *GeoTag) IsPointInArrondissement(arrondissement provider.ArrondissementFilter, coordinates []float64) bool {
	point := geo.NewPoint(coordinates[0], coordinates[1])
	switch arrondissement {
	case provider.Ahunstic:
		return geoTag.ahuntsic.Contains(point)
	case provider.MontRoyal:
		return geoTag.montRoyal.Contains(point)
	case provider.Rosemont:
		return geoTag.rosemont.Contains(point)
	case provider.TroisRivieres:
		return geoTag.mauricie.Contains(point)
	case provider.Villeray:
		return geoTag.villeray.Contains(point)
	}
	return true
}

// BuildGeoTag BuildGeoTag
func BuildGeoTag() *GeoTag {
	return &GeoTag{
		sweetSpot: buildPolygon("sweetspot"),
		ahuntsic:  buildPolygon("ahuntsic"),
		mauricie:  buildPolygon("mauricie"),
		montRoyal: buildPolygon("mont-royal"),
		rosemont:  buildPolygon("rosemont"),
		villeray:  buildPolygon("villeray"),
	}
}
