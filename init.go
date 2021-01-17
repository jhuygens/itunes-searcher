package itunes

import (
	"github.com/jgolang/config"
	searcher "github.com/jhuygens/searcher-engine"
)

var itunesSearcher = Searcher{}

func init() {
	searcher.RegisterSearcher(config.GetString("searchers.itunes"), itunesSearcher)
}
