package provider

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// CreateFormReader create form reader
func CreateFormReader(data map[string][]string) io.Reader {
	return strings.NewReader(CreateQueryString(data))
}

// CreateQueryString create query string
func CreateQueryString(data map[string][]string) string {
	form := url.Values{}
	for k, valueList := range data {
		for _, v := range valueList {
			form.Add(k, v)
		}
	}

	return form.Encode()
}

func ParsePrice(priceText string) int {
	priceReg := regexp.MustCompile("[^0-9]+")
	price, err := strconv.Atoi(priceReg.ReplaceAllString(priceText, ""))

	if err != nil {
		log.Fatal("could not parse price ", priceText)
	}

	return price
}

func ParseAddress(addressText string) string {
	streetNumberReg := regexp.MustCompile("\\d+")
	tempStreeName := streetNumberReg.Split(addressText, -1)
	streetName := removeNonAlphaNumeric(tempStreeName[len(tempStreeName)-1])
	streetNumber := removeNonAlphaNumeric(streetNumberReg.FindString(addressText))

	return fmt.Sprintf("%v %v", streetNumber, streetName)
}

func removeNonAlphaNumeric(stringToTreat string) string {
	nonAlphaNumericReg := regexp.MustCompile("[^A-zÀ-ú\\d\\s-]")

	return nonAlphaNumericReg.ReplaceAllString(stringToTreat, "")
}
