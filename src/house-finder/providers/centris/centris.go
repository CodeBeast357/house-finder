package centris

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	house "github.com/massintha/house-finder/src/house-finder/models"
	"github.com/massintha/house-finder/src/house-finder/providers/provider"
)

type getInscriptionsResponseResult struct {
	Count             int
	HTML              string
	InscNumberPerPage int
	Title             string
}

type getInscriptionsResponseD struct {
	Message   string
	Result    getInscriptionsResponseResult
	Succeeded bool
}

type getInscriptionsResponse struct {
	D getInscriptionsResponseD
}

// GetHouses gets houses
func GetHouses(arrondissement provider.ArrondissementFilter) []*house.House {
	c := colly.NewCollector(
		colly.AllowedDomains("centris.ca", "www.centris.ca"),
	)
	queryCollector := c.Clone()
	resultCollector := c.Clone()

	houses := make([]*house.House, 0)

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Something went wrong main:", err, r.StatusCode, string(r.Body))
	})
	queryCollector.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Something went wrong query:", err, r.StatusCode, string(r.Body))
	})
	resultCollector.OnError(func(r *colly.Response, err error) {
		log.Fatalln("Something went wrong results:", err, r.StatusCode, string(r.Body))
	})

	// Filters
	queryCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})

	resultCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json;charset=UTF-8")
	})

	resultCollector.OnResponse(func(r *colly.Response) {
		var response getInscriptionsResponse
		err := json.Unmarshal(r.Body, &response)
		if err != nil {
			log.Fatalln(err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(response.D.Result.HTML))
		if err != nil {
			log.Fatalln(err)
		}

		doc.Find(".shell").Each(func(_ int, propertyDoc *goquery.Selection) {
			house := getHouseItem(r, propertyDoc)
			house.Arrondissement = arrondissement
			houses = append(houses, house)
		})
		if len(houses) < response.D.Result.Count {
			triggerResultCollector(resultCollector, len(houses))
		}
	})

	c.Visit("https://www.centris.ca/")
	queryCollector.PostRaw("https://www.centris.ca/property/UpdateQuery", []byte(getPayload(arrondissement)))
	triggerResultCollector(resultCollector, 0)

	return houses
}

func triggerResultCollector(resultCollector *colly.Collector, startPosition int) {
	payload := fmt.Sprintf(`{"startPosition":%d}`, startPosition)
	resultCollector.PostRaw("https://www.centris.ca/Property/GetInscriptions", []byte(payload))
}

func getHouseItem(r *colly.Response, propertyDoc *goquery.Selection) *house.House {
	priceReg, _ := regexp.Compile("[^0-9]+")
	address := propertyDoc.Find(".address > div:nth-of-type(1)").Text()
	price, _ := strconv.Atoi(priceReg.ReplaceAllString(propertyDoc.Find(".price > span").Text(), ""))
	linkAttr, _ := propertyDoc.Find(".a-more-detail").Attr("href")
	link := r.Request.AbsoluteURL(linkAttr)
	thumbnailLink, _ := propertyDoc.Find(".thumbnail > a > img").Attr("src")

	return &house.House{
		Address:       address,
		Price:         price,
		Link:          link,
		ThumbnailLink: thumbnailLink,
		ProviderName:  "centris",
	}
}

func getPayload(arrondissement provider.ArrondissementFilter) string {
	plexFilters := `{
		"fieldId": "PropertyType",
		"value": "Plex",
		"fieldConditionId": "",
		"valueConditionId": "IsResidentialForSale"
	}, {
		"fieldId": "Plex",
		"value": "2X",
		"fieldConditionId": "IsPlex",
		"valueConditionId": ""
	}, {
		"fieldId": "Plex",
		"value": "3X",
		"fieldConditionId": "IsPlex",
		"valueConditionId": ""
	}`

	homeFilters := `{
		"fieldId": "PropertyType",
		"value": "SingleFamilyHome",
		"fieldConditionId": "",
		"valueConditionId": "IsResidential"
	}, {
		"fieldId": "PropertyType",
		"value": "SellCondo",
		"fieldConditionId": "",
		"valueConditionId": "IsResidential"
	}`

	arrondissementFilter := map[provider.ArrondissementFilter]map[string]string{
		provider.Ahunstic:      map[string]string{"matchType": "CityDistrictNeighbourhoodSearch", "id": "D;840", "text": "Quartier Ahuntsic Ouest, Montréal (Ahuntsic-Cartierville)", "propertyType": plexFilters},
		provider.MontRoyal:     map[string]string{"matchType": "CityDistrictNeighbourhoodSearch", "id": "D;842", "text": "Quartier Mile-End, Montréal (Le Plateau-Mont-Royal)", "propertyType": plexFilters},
		provider.Rosemont:      map[string]string{"matchType": "CityDistrictNeighbourhoodSearch", "id": "G;844", "text": "Quartier Petite Italie, Montréal (Rosemont/La Petite-Patrie)", "propertyType": plexFilters},
		provider.Villeray:      map[string]string{"matchType": "CityDistrictNeighbourhoodSearch", "id": "A;841", "text": "Quartier Villeray, Montréal (Villeray/Saint-Michel/Parc-Extension)", "propertyType": plexFilters},
		provider.TroisRivieres: map[string]string{"matchType": "CityDistrict", "id": "449", "text": "Trois-Rivières", "propertyType": homeFilters},
	}
	return fmt.Sprintf(`{
		"query": {
			"UseGeographyShapes": 0,
			"Filters": [{
				"MatchType": "%s",
				"Text": "%s",
				"Id": "%s"
			}],
			"FieldsValues": [{
				"fieldId": "%s",
				"value": "%s",
				"fieldConditionId": "",
				"valueConditionId": ""
			}, %s, {
				"fieldId": "Category",
				"value": "Residential",
				"fieldConditionId": "",
				"valueConditionId": ""
			}, {
				"fieldId": "SellingType",
				"value": "Sale",
				"fieldConditionId": "",
				"valueConditionId": ""
			}, {
				"fieldId": "LandArea",
				"value": "SquareFeet",
				"fieldConditionId": "IsLandArea",
				"valueConditionId": ""
			}, {
				"fieldId": "SalePrice",
				"value": %d,
				"fieldConditionId": "ForSale",
				"valueConditionId": ""
			}, {
				"fieldId": "SalePrice",
				"value": %d,
				"fieldConditionId": "ForSale",
				"valueConditionId": ""
			}]
		},
		"isHomePage": true
	}`,
		arrondissementFilter[arrondissement]["matchType"],
		arrondissementFilter[arrondissement]["text"],
		arrondissementFilter[arrondissement]["id"],
		arrondissementFilter[arrondissement]["matchType"],
		arrondissementFilter[arrondissement]["id"],
		arrondissementFilter[arrondissement]["propertyType"],
		provider.ArrondissementPriceFilter[arrondissement].Min,
		provider.ArrondissementPriceFilter[arrondissement].Max,
	)
}
