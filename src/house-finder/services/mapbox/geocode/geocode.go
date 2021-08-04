package geocode

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/go-querystring/query"
	house "github.com/massintha/house-finder/src/house-finder/models"
	"github.com/massintha/house-finder/src/house-finder/services/mapbox/mapbox"
)

// ErrorAPIUnauthorized indicates authorization failed
var ErrorAPIUnauthorized = errors.New("Mapbox API error unauthorized")

// ErrorAPILimitExceeded indicates the API limit has been exceeded
var ErrorAPILimitExceeded = errors.New("Mapbox API error api rate limit exceeded")

const (
	baseURL                 = "https://api.mapbox.com"
	apiName                 = "geocoding"
	apiVersion              = "v5"
	apiMode                 = "mapbox.places"
	statusRateLimitExceeded = 429
)

var token = os.Getenv("MAPBOX_TOKEN")

// APIMessage APIMessage
type APIMessage struct {
	Message string
}

// QueryRequest make a get with the provided query string and return the response if successful
func QueryRequest(query string, v *url.Values) (*http.Response, error) {
	// Add token to args
	v.Set("access_token", token)

	// Generate URL
	url := fmt.Sprintf("%s/%s", baseURL, query)

	fmt.Printf("URL: %s\n", url)

	// Create request object
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = v.Encode()

	// Create client instance
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// data, _ := httputil.DumpRequest(request, true)
	// fmt.Printf("Request: %s", string(data))
	// data, _ = httputil.DumpResponse(resp, false)
	// fmt.Printf("Response: %s", string(data))

	if resp.StatusCode == statusRateLimitExceeded {
		return nil, ErrorAPILimitExceeded
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrorAPIUnauthorized
	}

	return resp, nil
}

// QueryBase QueryBase
func QueryBase(query string, v *url.Values, inst interface{}) error {
	// Make request
	resp, err := QueryRequest(query, v)
	if err != nil && (resp == nil || resp.StatusCode != http.StatusBadRequest) {
		return err
	}
	defer resp.Body.Close()

	// Read body into buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Handle bad requests with messages
	if resp.StatusCode == http.StatusBadRequest {
		apiMessage := APIMessage{}
		messageErr := json.Unmarshal(body, &apiMessage)
		if messageErr == nil {
			return fmt.Errorf("api error: %s", apiMessage.Message)
		}
		return fmt.Errorf("Bad Request (400) - no message")
	}

	// Attempt to decode body into inst type
	err = json.Unmarshal(body, &inst)
	if err != nil {
		return err
	}

	return nil
}

// ForwardRequestOpts request options fo forward geocoding
type ForwardRequestOpts struct {
	BBox  mapbox.BoundingBox `url:"bbox,omitempty"`
	Limit uint               `url:"limit,omitempty"`
}

// ForwardResponse is the response from a forward geocode lookup
type ForwardResponse struct {
	*mapbox.FeatureCollection
}

// Forward Forward
func Forward(house *house.House, req *ForwardRequestOpts) (*ForwardResponse, error) {

	v, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	resp := ForwardResponse{}
	sanitizedPlace := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(house.Address, " ", "+"), "/", "+"), ";", "")

	queryString := fmt.Sprintf("%s/%s/%s/%s+%s+%s.json", apiName, apiVersion, apiMode, sanitizedPlace, house.Arrondissement, "Montreal")
	err = QueryBase(queryString, &v, &resp)

	return &resp, err
}
