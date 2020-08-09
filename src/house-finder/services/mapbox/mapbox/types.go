package mapbox

// Point Point
type Point []float64

// BoundingBox BoundingBox
type BoundingBox []float64

// Feature Feature
type Feature struct {
	Center Point `json:"center"`
}

// FeatureCollection FeatureCollection
type FeatureCollection struct {
	Features []Feature `json:"features"`
}
