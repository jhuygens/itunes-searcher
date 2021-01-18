package itunes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jgolang/config"
	"github.com/jgolang/consumer/rest"
	"github.com/jgolang/log"
	searcher "github.com/jhuygens/searcher-engine"
)

// Limit returns limit
var Limit = "20"

// Searcher searcher interface implement
type Searcher struct{}

// Search doc ...
func (s Searcher) Search(filter searcher.Filter) ([]searcher.Item, error) {
	if len(filter.Types) == 0 {
		filter.Types = append(filter.Types, "")
	}
	var items []searcher.Item
	var country string
	if len(filter.Country) > 1 {
		country = filter.Country[0].Value
	}
	for _, resource := range filter.Types {
		for _, name := range filter.Name {
			items = append(items, searchByName(name.Value, country, resource, Limit)...)
		}
		if len(filter.Name) == 0 {
			for _, artist := range filter.Artist {
				items = append(items, searchByName(artist.Value, country, "artist", Limit)...)
			}
			for _, album := range filter.Album {
				items = append(items, searchByName(album.Value, country, "album", Limit)...)
			}
		}
	}
	return items, nil
}

func searchByName(term, country, resource, limit string) []searcher.Item {
	media := media[resource]
	queryParams := make(map[string]string)
	queryParams["term"] = term
	queryParams["country"] = country
	queryParams["limit"] = limit
	queryParams["media"] = media
	// queryParams["entity"] = mediaTypeEntities[media][resource]
	queryParams["attribute"] = mediaTypeAtributes[media][resource]
	result, err := searchItunesServiceItems(queryParams)
	if err != nil {
		log.Error(err)
	}
	var items []searcher.Item
	items = append(items, result...)
	return items
}

func searchItunesServiceItems(queryParams map[string]string) ([]searcher.Item, error) {
	resquest := rest.RequestInfo{
		Method:      http.MethodGet,
		Endpoint:    config.GetString("integrations.itunes.endpoint"),
		Timeout:     time.Duration(config.GetInt("integrations.itunes.timeout")) * time.Second,
		QueryParams: queryParams,
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
		ratingAvg, _ := strconv.ParseFloat(result.ContentAdvisoryRating, 64)
		items = append(
			items,
			searcher.Item{
				Type:    result.Kind,
				Library: config.GetString("searchers.itunes"),
				Name:    result.TrackName,
				Artwork: result.ArtworkURL100,
				Info: searcher.Info{
					PreviewURL:  result.PreviewURL,
					Title:       result.TrackName,
					Collection:  result.CollectionName,
					Artist:      result.ArtistName,
					Languages:   nil,
					RatingAvg:   ratingAvg,
					Genres:      []string{result.PrimaryGenreName},
					Description: result.ShortDescription,
					MoreInfo:    result.LongDescription,
					ReleaseDate: result.ReleaseDate,
					Country:     result.Country,
					Price:       result.TrackPrice,
					RentalPrice: result.TrackHdRentalPrice,
					Currency:    result.Currency,
					URL:         result.TrackViewURL,
				},
			},
		)
	}
	return items, nil
}
