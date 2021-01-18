package itunes

// ResponseData itunes
type ResponseData struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

// Result itunes request
type Result struct {
	WrapperType             string  `json:"wrapperType"`
	Kind                    string  `json:"kind"`
	TrackID                 int     `json:"trackId"`
	ArtistName              string  `json:"artistName"`
	TrackName               string  `json:"trackName"`
	TrackCensoredName       string  `json:"trackCensoredName"`
	TrackViewURL            string  `json:"trackViewUrl"`
	PreviewURL              string  `json:"previewUrl"`
	ArtworkURL30            string  `json:"artworkUrl30"`
	ArtworkURL60            string  `json:"artworkUrl60"`
	ArtworkURL100           string  `json:"artworkUrl100"`
	CollectionPrice         float64 `json:"collectionPrice"`
	TrackPrice              float64 `json:"trackPrice"`
	TrackRentalPrice        float64 `json:"trackRentalPrice"`
	CollectionHdPrice       float64 `json:"collectionHdPrice"`
	TrackHdPrice            float64 `json:"trackHdPrice"`
	TrackHdRentalPrice      float64 `json:"trackHdRentalPrice"`
	ReleaseDate             string  `json:"releaseDate"`
	CollectionExplicitness  string  `json:"collectionExplicitness"`
	TrackExplicitness       string  `json:"trackExplicitness"`
	TrackTimeMillis         int     `json:"trackTimeMillis"`
	Country                 string  `json:"country"`
	Currency                string  `json:"currency"`
	PrimaryGenreName        string  `json:"primaryGenreName"`
	ContentAdvisoryRating   string  `json:"contentAdvisoryRating"`
	ShortDescription        string  `json:"shortDescription,omitempty"`
	LongDescription         string  `json:"longDescription"`
	CollectionID            int     `json:"collectionId,omitempty"`
	CollectionName          string  `json:"collectionName,omitempty"`
	CollectionCensoredName  string  `json:"collectionCensoredName,omitempty"`
	CollectionArtistID      int     `json:"collectionArtistId,omitempty"`
	CollectionArtistViewURL string  `json:"collectionArtistViewUrl,omitempty"`
	CollectionViewURL       string  `json:"collectionViewUrl,omitempty"`
	DiscCount               int     `json:"discCount,omitempty"`
	DiscNumber              int     `json:"discNumber,omitempty"`
	TrackCount              int     `json:"trackCount,omitempty"`
	TrackNumber             int     `json:"trackNumber,omitempty"`
	HasITunesExtras         bool    `json:"hasITunesExtras,omitempty"`
	ArtistID                int     `json:"artistId,omitempty"`
	ArtistViewURL           string  `json:"artistViewUrl,omitempty"`
}
