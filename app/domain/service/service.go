package service

// package responsible for manage the business rule

import (
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/cache" // package responsible for save cache data
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/swapi" // package responsible for provide swapi data
	"github.com/SousaGLucas/swsearch/app/domain/entities/swdata"    // package responsible for mirror database data
)

type Service interface {
	GetData(searchTerm string) (swdata.Data, error) // resposible for get data that each route in the swapi
	GetCache() (swdata.Cache, error)                // responsible for get cache data
	ClearCache() error                              // responsible for call clear cache methode
}

// mirror that database data
type Result swdata.Data

// resposible for get data that each route in the swapi
func (result Result) GetData(searchTerm string) (swdata.Data, error) {
	emptyData := swdata.Data{} // empty variable to return in the error cases
	cacheData := cache.Data{}
	response := swapi.Result{}

	// check if the term already serached
	if searched := cacheData.CheckTerm(searchTerm); searched != nil {
		return emptyData, searched
	}

	// get response data for the database
	data, err := response.Search(searchTerm)

	if err != nil {
		return emptyData, err
	}

	// adding term in the cache
	pushErr := cacheData.Push(searchTerm)

	if pushErr != nil {
		return emptyData, pushErr
	}

	return data, nil
}

// responsible for call clear cache methode
func (result Result) ClearCache() error {
	cacheData := cache.Data{}

	err := cacheData.Clear() // call clear cache methode

	if err != nil {
		return err
	}

	return nil
}

// responsible for get cache data
func (result Result) GetCache() (swdata.Cache, error) {
	emptyCache := swdata.Cache{} // empty variable to return in the error cases
	cacheData := cache.Data{}

	cache, err := cacheData.GetCache() // call clear cache methode

	if err != nil {
		return emptyCache, err
	}

	return cache, nil
}
