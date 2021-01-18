package itunes

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
