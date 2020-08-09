package provider

// ArrondissementFilter ArrondissementFilter
type ArrondissementFilter string

const (
	// Ahunstic Ahunstic
	Ahunstic ArrondissementFilter = "Ahunstic"
	// MontRoyal MontRoyal
	MontRoyal ArrondissementFilter = "MontRoyal"
	// Rosemont Rosemont
	Rosemont ArrondissementFilter = "Rosemont"
	// Villeray Villeray
	Villeray ArrondissementFilter = "Villeray"

	// TroisRivieres TroisRivieres
	TroisRivieres ArrondissementFilter = "Trois-Rivieres"
)

type priceFilter struct {
	Min int
	Max int
}

// ArrondissementPriceFilter ArrondissementPriceFilter
var ArrondissementPriceFilter = map[ArrondissementFilter]priceFilter{
	Ahunstic:      priceFilter{Min: 0, Max: 1000000},
	MontRoyal:     priceFilter{Min: 0, Max: 1000000},
	Rosemont:      priceFilter{Min: 0, Max: 1000000},
	Villeray:      priceFilter{Min: 0, Max: 1000000},
	TroisRivieres: priceFilter{Min: 175000, Max: 400000},
}

// ArrondissementBBox ArrondissementBBox
var ArrondissementBBox = map[ArrondissementFilter][]float64{
	Ahunstic:      []float64{-73.815958, 45.500997, -73.562176, 45.624338},
	MontRoyal:     []float64{-73.815958, 45.500997, -73.562176, 45.624338},
	Rosemont:      []float64{-73.815958, 45.500997, -73.562176, 45.624338},
	Villeray:      []float64{-73.815958, 45.500997, -73.562176, 45.624338},
	TroisRivieres: []float64{-72.779418, 46.252823, -72.408782, 46.469151},
}
