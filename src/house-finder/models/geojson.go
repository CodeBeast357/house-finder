package models

// GeoJSON type
type GeoJSON struct {
	Type        string        `json:"type" db:"type"`
	Coordinates [][][]float64 `json:"coordinates" db:"coordinates"`
}
