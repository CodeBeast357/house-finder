package models

import "github.com/massintha/house-finder/src/house-finder/providers/provider"

// House type
type House struct {
	ID               int                           `json:"id" db:"id"`
	Address          string                        `json:"address" db:"address"`
	Arrondissement   provider.ArrondissementFilter `json:"arrondissement" db:"arrondissement"`
	Price            int                           `json:"price" db:"price"`
	Link             string                        `json:"link" db:"link"`
	ThumbnailLink    string                        `json:"thumbnailLink" db:"thumbnail_link"`
	ProviderName     string                        `json:"providerName" db:"provider_name"`
	Latitude         float64                       `json:"latitude" db:"latitude"`
	Longitude        float64                       `json:"longitude" db:"longitude"`
	IsBlackListed    bool                          `json:"isBlackListed" db:"is_black_listed"`
	IsFavorite       bool                          `json:"isFavorite" db:"is_favorite"`
	IsInSweetSpot    bool                          `json:"isInSweetSpot" db:"is_in_sweet_spot"`
	CreationDatetime float64                       `json:"creationDatetime" db:"creation_datetime"`
}

// SetCoordinates SetCoordinates
func (house *House) SetCoordinates(coordinates []float64) {
	house.Longitude = coordinates[0]
	house.Latitude = coordinates[1]
}

// SetIsInSweetSpot SetIsInSweetSpot
func (house *House) SetIsInSweetSpot(isInSweetSpot bool) {
	house.IsInSweetSpot = isInSweetSpot
}
