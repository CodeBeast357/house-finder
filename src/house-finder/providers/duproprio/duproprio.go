package duproprio

import (
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	house "github.com/massintha/house-finder/src/house-finder/models"
	"github.com/massintha/house-finder/src/house-finder/providers/provider"
)

// GetHouses gets houses
func GetHouses(arrondissement provider.ArrondissementFilter) []*house.House {
	resultCollector := colly.NewCollector(
		colly.AllowedDomains("duproprio.com", "www.duproprio.com"),
	)
	houses := make([]*house.House, 0)

	// Error handling
	resultCollector.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Something went wrong results:", err, r.StatusCode, string(r.Body))
	})

	resultCollector.OnHTML(".search-results-listings-list__item", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("id"), "listing") && !e.DOM.HasClass("is-sold") {
			house := getHouseItem(e)
			house.Arrondissement = arrondissement
			houses = append(houses, house)
		}
	})

	resultCollector.OnHTML(".search-results-listings-header__properties-found__number", func(e *colly.HTMLElement) {
		params, _ := url.ParseQuery(e.Request.URL.RawQuery)
		currentPageNumber, _ := strconv.Atoi(params.Get("pageNumber"))
		numberOfProperties, _ := strconv.Atoi(e.Text)
		if len(houses) < numberOfProperties {
			triggerResultCollector(arrondissement, resultCollector)(currentPageNumber + 1)
		}
	})

	triggerResultCollector(arrondissement, resultCollector)(1)

	return houses
}

func triggerResultCollector(arrondissement provider.ArrondissementFilter, resultCollector *colly.Collector) func(pageNumber int) {
	arrondissementFilter := map[provider.ArrondissementFilter]map[string][]string{
		provider.Ahunstic:      {"cities": {"1893"}, "type": {"multiplex"}, "subtype": {"52", "53"}},
		provider.MontRoyal:     {"cities": {"1887"}, "type": {"multiplex"}, "subtype": {"52", "53"}},
		provider.Rosemont:      {"cities": {"1889"}, "type": {"multiplex"}, "subtype": {"52", "53"}},
		provider.Villeray:      {"cities": {"1892"}, "type": {"multiplex"}, "subtype": {"52", "53"}},
		provider.TroisRivieres: {"cities": {"1037"}, "type": {"house", "condo"}, "subtype": {"1", "2", "4", "5", "6", "7", "9", "10", "11", "13", "15", "17", "19", "21", "97", "99", "100", "3", "12", "14", "105"}},
	}

	return func(pageNumber int) {
		queryString := provider.CreateQueryString(map[string][]string{
			"search":        {"true"},
			"cities[]":      arrondissementFilter[arrondissement]["cities"],
			"min_price":     {strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Min)},
			"max_price":     {strconv.Itoa(provider.ArrondissementPriceFilter[arrondissement].Max)},
			"type[]":        arrondissementFilter[arrondissement]["type"],
			"subtype[]":     arrondissementFilter[arrondissement]["subtype"],
			"is_for_sale":   {"1"},
			"with_builders": {"1"},
			"parent":        {"1"},
			"pageNumber":    {strconv.Itoa(pageNumber)},
			"sort":          {"-published_at"},
		})
		base, _ := url.Parse("https://duproprio.com/fr/rechercher/liste")
		base.RawQuery = queryString
		resultCollector.Visit(base.String())
	}
}

func getHouseItem(e *colly.HTMLElement) *house.House {
	address := provider.ParseAddress(e.ChildText(".search-results-listings-list__item-description__address"))
	price := provider.ParsePrice(e.ChildText(".search-results-listings-list__item-description__price"))
	link := e.Request.AbsoluteURL(e.ChildAttr(".search-results-listings-list__item-image-link ", "href"))
	thumbnailLink := e.ChildAttr(".search-results-listings-list__item-photo", "src")
	return &house.House{
		Address:       address,
		Price:         price,
		Link:          link,
		ThumbnailLink: thumbnailLink,
		ProviderName:  "duproprio",
	}
}
