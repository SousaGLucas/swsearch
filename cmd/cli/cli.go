package cli

import (
	"flag"
	"fmt"

	"github.com/SousaGLucas/swsearch/app/gateway/controller" // responsible for accessing the domain methods
)

func Execute() {
	searchTerm := flag.String("s", "", "fetches data according to entered search term") // search data flag
	clearCache1 := flag.Bool("c", false, "clear cache")                                 // clear cache flag
	clearCache2 := flag.Bool("clear", false, "clear cache")                             // clear cache flag
	showCache := flag.Bool("sc", false, "show cache")                                   // show cache flag
	flag.Bool("h", false, "show help")                                                  // show help flag
	flag.Bool("help", false, "show help")                                               // show help flag
	flag.Parse()

	if *searchTerm != "" {
		data, err := controller.Search(*searchTerm) // call search data function

		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Println(data)
		return
	}

	if *clearCache1 || *clearCache2 {
		err := controller.ClearChache() // call clear cache function

		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("cache deleted\n")
		return
	}

	if *showCache {
		cache, err := controller.GetCache() // call get cache function

		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("Cache: %v\n", cache)
		return
	}

	// help message

	fmt.Printf("\nSWSEARCH\n")
	fmt.Printf("App for searching data from the Star Wars universe.\n\n")
	fmt.Printf("Version: development\n\n")

	fmt.Printf("Usage:\n\n")
	fmt.Printf("   -c, -clear		clear cache\n")
	fmt.Printf("   -h, -help		show help message\n")
	fmt.Printf("   -s [string]		fetches data according to entered search term\n")
	fmt.Printf("   -sc			show cache\n")

	fmt.Printf("\n\n")
	return
}
