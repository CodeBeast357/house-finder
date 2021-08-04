package remax

import (
	"log"
	"strconv"
	"strings"

	colly "github.com/gocolly/colly"
	house "github.com/massintha/house-finder/src/house-finder/models"
	"github.com/massintha/house-finder/src/house-finder/providers/provider"
)

// GetHouses gets houses
func GetHouses(arrondissement provider.ArrondissementFilter) []*house.House {
	c := colly.NewCollector(
		colly.AllowedDomains("remax-quebec.com", "www.remax-quebec.com"),
	)

	houses := make([]*house.House, 0)

	c.OnHTML(".property-entry", func(e *colly.HTMLElement) {
		address := strings.Split(e.ChildText(".property-address > h2"), ",")[0]
		price := provider.ParsePrice(e.ChildText(".property-price"))
		link := e.Request.AbsoluteURL(e.ChildAttr(".property-details", "href"))
		thumbnailLink := e.ChildAttr(".property-thumbnail > img", "src")
		house := &house.House{
			Address:        address,
			Price:          price,
			Arrondissement: arrondissement,
			Link:           link,
			ThumbnailLink:  thumbnailLink,
			ProviderName:   "remax",
		}

		houses = append(houses, house)
	})

	c.OnHTML(".pagination > li > a[aria-label='Suivant']", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	arrondissementFilter := map[provider.ArrondissementFilter]map[string][]string{
		provider.Ahunstic:      {"cities": {"66506"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.MontRoyal:     {"cities": {"66508"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.Rosemont:      {"cities": {"6651171"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.Villeray:      {"cities": {"66507"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.TroisRivieres: {"cities": {"37067", "37055", "37070", "37075", "37050", "37060"}, "regions": {"4"}, "categories": {"plex"}, "genres": {"1", "2", "3"}},
	}

	formData := provider.CreateFormReader(map[string][]string{
		"mode":                   {"criterias"},
		"order":                  {"date_desc"},
		"query":                  {""},
		"categorie":              arrondissementFilter[arrondissement]["categories"],
		"selectItemcategorie":    arrondissementFilter[arrondissement]["categories"],
		"genres":                 arrondissementFilter[arrondissement]["genres"],
		"selectItemgenres":       arrondissementFilter[arrondissement]["genres"],
		"regionIds":              arrondissementFilter[arrondissement]["regions"],
		"selectItemregionIds":    arrondissementFilter[arrondissement]["regions"],
		"cityIds":                arrondissementFilter[arrondissement]["cities"],
		"selectItemcityIds":      arrondissementFilter[arrondissement]["cities"],
		"minPrice":               {strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Min)},
		"selectItemminPrice":     {strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Min)},
		"maxPrice":               {strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Max)},
		"selectItemmaxPrice":     {strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Max)},
		"transacTypes":           {"vente"},
		"selectItemtransacTypes": {"vente"},
	})

	c.Request("POST", "https://www.remax-quebec.com/en/recherche/plex/resultats.rmx", formData, nil, nil)

	return houses
}
