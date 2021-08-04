package provider

import (
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
	priceReg, _ := regexp.Compile("[^0-9]+")
	price, err := strconv.Atoi(priceReg.ReplaceAllString(priceText, ""))

	if err != nil {
		log.Fatal("could not parse price ", priceText)
	}

	return price
}
