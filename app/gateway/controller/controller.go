package controller

// responsible for accessing the domain methods

import (
	"encoding/json"
	"errors"

	// package responsible for mirror database data
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/cache"
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/swapi"
	"github.com/SousaGLucas/swsearch/app/domain/service" // package responsible for manage the business rule
	"github.com/SousaGLucas/swsearch/log"                // package responsible for logging system errors
)

// resposible for get survey data
func Search(searchTerm string) (string, error) {
	response := service.Result{}
	cacheData := cache.Data{}
	responseData := swapi.Result{}

	data, err := response.GetData(searchTerm, responseData, cacheData) // get survey data

	if err != nil {
		return "[]", err
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.SetLog(err)
		return "[]", errors.New("error reading data")
	}

	return string(jsonData), nil
}

// responsible for clear cache data
func ClearChache() error {
	response := service.Result{}
	cacheData := cache.Data{}

	err := response.ClearCache(cacheData) // clear cache data

	if err != nil {
		return err
	}

	return nil
}

// responsible for get cache data
func GetCache() (string, error) {
	response := service.Result{}
	cacheData := cache.Data{}

	cache, err := response.GetCache(cacheData) // get current cache data

	if err != nil {
		return "[]", err
	}

	jsonCache, err := json.Marshal(cache)

	if err != nil {
		log.SetLog(err)
		return "[]", errors.New("error reading cache")
	}

	return string(jsonCache), nil
}
