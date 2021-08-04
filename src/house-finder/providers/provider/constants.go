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
	Ahunstic:      {Min: 0, Max: 1500000},
	MontRoyal:     {Min: 0, Max: 1500000},
	Rosemont:      {Min: 0, Max: 1500000},
	Villeray:      {Min: 0, Max: 1500000},
	TroisRivieres: {Min: 175000, Max: 400000},
}

// ArrondissementBBox ArrondissementBBox
var ArrondissementBBox = map[ArrondissementFilter][]float64{
	Ahunstic:      {-73.972902, 45.410076, -73.474295, 45.70479},
	MontRoyal:     {-73.972902, 45.410076, -73.474295, 45.70479},
	Rosemont:      {-73.972902, 45.410076, -73.474295, 45.70479},
	Villeray:      {-73.972902, 45.410076, -73.474295, 45.70479},
	TroisRivieres: {-72.779418, 46.252823, -72.408782, 46.469151},
}
