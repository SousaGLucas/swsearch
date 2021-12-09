package main

import (
	"fmt"

	"github.com/SousaGLucas/swsearch/domain/dataprovider/cache" // package responsible for save cache data
	"github.com/SousaGLucas/swsearch/domain/service"            // package responsible for manage the business rule
)

func main() {
	fmt.Printf("CONSULATA 1: 'Sky'\n\n")
	search("Sky") // search function call

	fmt.Printf("CONSULATA 2: 'r2'\n\n")
	search("r2") // search function call

	fmt.Printf("CONSULATA 3: 'Sky'\n\n")
	search("Sky") // search function call

	clearChache() // clear cache
	fmt.Printf("\nCACHE DELETADO!\n\n")

	fmt.Printf("CONSULATA 4: 'Sky'\n\n")
	search("Sky") // search function call
}

// resposible for get survey data --> **test function
func search(searchTerm string) {
	var response service.Service
	response = service.Result{}

	data, err := response.GetData(searchTerm) // get survey data

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	fmt.Printf("Responsta: %v\n\n", data) // print response

	cacheData := cache.GetCache()          // get current cache data
	fmt.Printf("CACHE: %v\n\n", cacheData) // print current cache data
	return
}

// responsible for clear cache data --> **test function
func clearChache() {
	var response service.Service
	response = service.Result{}

	response.ClearCache() // clear cache data
	return
}
