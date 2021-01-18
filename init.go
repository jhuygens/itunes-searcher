package itunes

import (
	"github.com/jgolang/config"
	"github.com/jgolang/log"
	searcher "github.com/jhuygens/searcher-engine"
)

var itunesSearcher = Searcher{}

func init() {
	name := config.GetString("searchers.itunes")
	err := searcher.RegisterSearcher(name, itunesSearcher)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("Searcher %v has been register", name)
}
