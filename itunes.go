package itunes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jgolang/config"
	"github.com/jgolang/consumer/rest"
	searcher "github.com/jhuygens/searcher-engine"
)

// Searcher searcher interface implement
type Searcher struct{}

// Search doc ...
func (s Searcher) Search(filter searcher.Filter) ([]searcher.Item, error) {
	var items []searcher.Item
	for _, resourceType := range filter.Types {
		for _, name := range filter.Name {
			query := make(map[string]string)
			query["limit"] = "20"
			query["name"] = name.Value
			query["entity"] = resourceType
			if len(filter.Country) > 1 {
				query["country"] = filter.Country[0].Value
			}
			result, err := searchItunesServiceItems(resourceType, query)
			if err != nil {
				return nil, err
			}
			items = append(items, result...)
		}
	}
	return items, nil
}

func searchItunesServiceItems(typeItem string, q map[string]string) ([]searcher.Item, error) {
	resquest := rest.RequestInfo{
		Method:   http.MethodGet,
		Endpoint: config.GetString("integrations.itunes.endpoint"),
		Timeout:  time.Duration(config.GetInt("integrations.itunes.timeout")) * time.Second,
		Query:    q,
	}
	var responseData ResponseData
	response, err := rest.ConsumeRestService(resquest, &responseData)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: status code %v", response.StatusCode)
	}
	var items []searcher.Item
	for _, result := range responseData.Results {
		items = append(
			items,
			searcher.Item{
				Type:    typeItem,
				Library: "itunes",
				Name:    result.TrackName,
				Artwork: result.ArtworkURL100,
				// Info:    result,
			},
		)
	}
	return items, nil
}
