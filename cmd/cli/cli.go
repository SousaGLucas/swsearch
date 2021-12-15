package cli

import (
	"flag"

	"github.com/SousaGLucas/swsearch/app/gateway/controller" // responsible for accessing the domain methods
)

func Execute() {
	searchTerm := flag.String("s", "", "fetches data according to entered search term") // search data flag
	clearCache := flag.Bool("c", false, "clear cache")                                  // clear cache flag
	showCache := flag.Bool("sc", false, "show cache")                                   // show cache flag
	flag.Parse()

	if *searchTerm != "" {
		controller.Search(*searchTerm) // call search data function
		return
	}

	if *clearCache {
		controller.ClearChache() // call clear cache function
		return
	}

	if *showCache {
		controller.GetCache() // call get cache function
		return
	}

	flag.CommandLine.Usage() // show help message
}
