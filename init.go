package itunes

import (
	"log"

	"github.com/jgolang/config"
	searcher "github.com/jhuygens/searcher-engine"
)

var itunesSearcher = Searcher{}

func init() {
	err := searcher.RegisterSearcher(config.GetString("searchers.itunes"), itunesSearcher)
	if err != nil {
		log.Fatal(err)
		return
	}
}
