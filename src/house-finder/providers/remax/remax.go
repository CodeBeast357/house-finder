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
		price, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(e.ChildText(".property-price"), "$", ""), ",", ""))
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
		provider.Ahunstic:      map[string][]string{"cities": {"66506"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.MontRoyal:     map[string][]string{"cities": {"66508"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.Rosemont:      map[string][]string{"cities": {"6651171"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.Villeray:      map[string][]string{"cities": {"66507"}, "regions": {"20"}, "categories": {"plex"}, "genres": {"2", "1"}},
		provider.TroisRivieres: map[string][]string{"cities": {"37067", "37055", "37070", "37075", "37050", "37060"}, "regions": {"4"}, "categories": {"plex"}, "genres": {"1", "2", "3"}},
	}

	formData := provider.CreateFormReader(map[string][]string{
		"mode":                   []string{"criterias"},
		"order":                  []string{"date_desc"},
		"query":                  []string{""},
		"categorie":              arrondissementFilter[arrondissement]["categories"],
		"selectItemcategorie":    arrondissementFilter[arrondissement]["categories"],
		"genres":                 arrondissementFilter[arrondissement]["genres"],
		"selectItemgenres":       arrondissementFilter[arrondissement]["genres"],
		"regionIds":              arrondissementFilter[arrondissement]["regions"],
		"selectItemregionIds":    arrondissementFilter[arrondissement]["regions"],
		"cityIds":                arrondissementFilter[arrondissement]["cities"],
		"selectItemcityIds":      arrondissementFilter[arrondissement]["cities"],
		"minPrice":               []string{strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Min)},
		"selectItemminPrice":     []string{strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Min)},
		"maxPrice":               []string{strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Max)},
		"selectItemmaxPrice":     []string{strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Max)},
		"transacTypes":           []string{"vente"},
		"selectItemtransacTypes": []string{"vente"},
	})

	c.Request("POST", "https://www.remax-quebec.com/en/recherche/plex/resultats.rmx", formData, nil, nil)

	return houses
}
