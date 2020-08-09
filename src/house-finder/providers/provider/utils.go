package provider

import (
	"io"
	"net/url"
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
