package itunes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jgolang/config"
	"github.com/jgolang/consumer/rest"
	"github.com/jgolang/log"
	searcher "github.com/jhuygens/searcher-engine"
)

// Searcher searcher interface implement
type Searcher struct{}

// Limit returns limit
var Limit = "10"

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
	}
	// No aplica el flitro
	// for _, artist := range filter.Artist {
	// 	items = append(items, searchByName(artist.Value, country, "artist", Limit)...)
	// }
	// for _, album := range filter.Album {
	// 	items = append(items, searchByName(album.Value, country, "album", Limit)...)
	// }

	return items, nil
}

func searchItunesServiceItems(q map[string]string) ([]searcher.Item, error) {
	resquest := rest.RequestInfo{
		Method:      http.MethodGet,
		Endpoint:    config.GetString("integrations.itunes.endpoint"),
		Timeout:     time.Duration(config.GetInt("integrations.itunes.timeout")) * time.Second,
		QueryParams: q,
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
				Type:    result.Kind,
				Library: "itunes",
				Name:    result.TrackName,
				Artwork: result.ArtworkURL100,
				// Info:    result,
			},
		)
	}
	return items, nil
}

func searchByName(term, country, resource, limit string) []searcher.Item {
	media := media[resource]
	query := make(map[string]string)
	query["term"] = term
	query["country"] = country
	query["limit"] = limit
	query["media"] = media
	query["entity"] = mediaTypeEntities[media][resource]
	query["attribute"] = mediaTypeAtributes[media][resource]
	result, err := searchItunesServiceItems(query)
	if err != nil {
		log.Error(err)
	}
	var items []searcher.Item
	items = append(items, result...)
	return items
}

var media = map[string]string{
	"album":  "music",
	"artist": "music",
	// "playlist": "", // N/A
	"track":   "music",
	"show":    "tvShow",
	"movie":   "movie",
	"podcast": "podcast",
	// "people":  "", // N/A
	"episode": "tvShow",
}

var mediaTypeAtributes = map[string]map[string]string{
	"music":   musicAttributes,
	"movie":   movieAttributes,
	"podcast": podcastAttributes,
	"tvShow":  tvShowAttributes,
}

var movieAttributes = map[string]string{
	"artist":  "artistTerm",
	"track":   "shortFilmTerm",
	"movie":   "movieTerm",
	"people":  "actorTerm",
	"episode": "shortFilmTerm",
}

var musicAttributes = map[string]string{
	"album":  "albumTerm",
	"artist": "artistTerm",
	"track":  "songTerm",
	"people": "composerTerm",
}

var podcastAttributes = map[string]string{
	"album":   "keywordsTerm",
	"artist":  "artistTerm",
	"podcast": "titleTerm",
	"people":  "authorTerm",
}

var tvShowAttributes = map[string]string{
	"album":    "tvSeasonTerm",
	"playlist": "tvSeasonTerm",
	"track":    "tvEpisodeTerm",
	"show":     "showTerm",
	"movie":    "movieTerm",
	"episode":  "tvEpisodeTerm",
}

var mediaTypeEntities = map[string]map[string]string{
	"music":   musicEntities,
	"movie":   movieEntities,
	"podcast": podcastEntities,
	"tvShow":  tvShowEntities,
}

var musicEntities = map[string]string{
	"album":  "album",
	"artist": "musicArtist",
	"track":  "song",
	"people": "musicArtist",
}

var movieEntities = map[string]string{
	"artist":  "movieArtist",
	"track":   "movie",
	"movie":   "movie",
	"people":  "movieArtist",
	"episode": "movie",
}

var podcastEntities = map[string]string{
	"album":   "podcast",
	"artist":  "podcastAuthor",
	"podcast": "podcast",
	"people":  "podcastAuthor",
}

var tvShowEntities = map[string]string{
	"album":    "tvSeason",
	"playlist": "tvSeason",
	"track":    "tvEpisode",
	"show":     "tvEpisode",
	"movie":    "tvEpisode",
	"episode":  "tvEpisode",
}
