package store

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/massintha/house-finder/src/house-finder/models"
)

const dataPath = "/data/metastore.db"
const bucketName = "house-finder"

// Store store
type Store struct {
	db *sqlx.DB
}

// SetupStore create db connection
func SetupStore() (*Store, error) {
	const (
		host     = "house_finder_db"
		port     = 5432
		user     = "root"
		password = "root"
		dbname   = "house_finder"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	return &Store{db}, nil
}

// SaveHouseList saves into the DB the list of given houses
func (store *Store) SaveHouseList(houseList []*models.House) error {
	insertQuery := `
        INSERT INTO houses(address, price, arrondissement, link, thumbnail_link, provider_name, coordinates, is_in_sweet_spot)
        VALUES (:address, :price, :arrondissement, :link, :thumbnail_link, :provider_name, ST_Point(:longitude, :latitude), :is_in_sweet_spot)
        ON CONFLICT DO NOTHING;
    `

	tx := store.db.MustBegin()
	for _, house := range houseList {
		tx.NamedExec(insertQuery, house)
	}
	return tx.Commit()
}

// GetAllHouses to get all the houses
func (store *Store) GetAllHouses() ([]models.House, error) {
	query := `SELECT
        id,
        address,
        price,
        arrondissement,
        link,
        thumbnail_link,
        provider_name,
        ST_X(coordinates) as latitude,
		ST_Y(coordinates) as longitude,
		is_black_listed,
        is_in_sweet_spot,
		is_favorite,
		extract(epoch from creation_datetime) * 1000 as creation_datetime
		FROM houses
		ORDER BY id;
    `
	houses := []models.House{}
	err := store.db.Select(&houses, query)
	return houses, err
}

// GetBoundingBox to get all the houses
func (store *Store) GetBoundingBox() (models.GeoJSON, error) {
	row := store.db.QueryRow(`SELECT ST_AsGeoJSON(ST_Extent(coordinates)) as geometry FROM houses;`)

	var geom string
	err := row.Scan(&geom)
	geometry := models.GeoJSON{}
	json.Unmarshal([]byte(geom), &geometry)
	return geometry, err
}

// UpdateHouse to get all the houses
func (store *Store) UpdateHouse(house models.House) error {
	updateQuery := `
        UPDATE houses SET is_favorite = :is_favorite, is_black_listed = :is_black_listed WHERE id = :id;
    `

	_, err := store.db.NamedExec(updateQuery, house)
	return err
}
