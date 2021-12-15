package controller

// responsible for accessing the domain methods

import (
	"fmt"

	"github.com/SousaGLucas/swsearch/app/domain/service" // package responsible for manage the business rule
)

// resposible for get survey data
func Search(searchTerm string) {
	var response service.Service
	response = service.Result{}

	data, err := response.GetData(searchTerm) // get survey data

	if err != nil {
		fmt.Printf("ERROR: %v\n", err) // print survey error
		return
	}

	fmt.Printf("Responsta: %v\n\n", data) // print response
	return
}

// responsible for clear cache data
func ClearChache() error {
	var response service.Service
	response = service.Result{}

	err := response.ClearCache() // clear cache data

	if err != nil {
		return err
	}

	fmt.Printf("cache deleted\n")
	return nil
}

func GetCache() error {
	var response service.Service
	response = service.Result{}

	cacheData, err := response.GetCache() // get current cache data

	if err != nil {
		return err
	}

	fmt.Printf("CACHE: %v\n", cacheData) // print current cache data
	return nil
}
