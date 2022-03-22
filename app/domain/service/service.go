package service

// package responsible for manage the business rule

import (
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/cache" // package responsible for save cache data
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/swapi" // package responsible for provide swapi data
	"github.com/SousaGLucas/swsearch/app/domain/entities/swdata"    // package responsible for mirror database data
)

type Service interface {
	GetData(searchTerm string, response swapi.Swapi, cacheData cache.Cache) (swdata.Data, error) // resposible for get data that each route in the swapi
	GetCache(cacheData cache.Cache) (swdata.Cache, error)                                        // responsible for get cache data
	ClearCache(cacheData cache.Cache) error                                                      // responsible for call clear cache methode
}

// mirror that database data
type Result swdata.Data

// resposible for get data that each route in the swapi
func (result Result) GetData(searchTerm string, response swapi.Swapi, cacheData cache.Cache) (swdata.Data, error) {
	// check if the term already serached
	if searched := cacheData.CheckTerm(searchTerm); searched != nil {
		return swdata.Data{}, searched
	}

	// get response data for the database
	data, err := response.Search(searchTerm)

	if err != nil {
		return swdata.Data{}, err
	}

	// adding term in the cache
	pushErr := cacheData.Push(searchTerm)

	if pushErr != nil {
		return swdata.Data{}, pushErr
	}

	return data, nil
}

// responsible for call clear cache methode
func (result Result) ClearCache(cacheData cache.Cache) error {
	err := cacheData.Clear() // call clear cache methode

	if err != nil {
		return err
	}

	return nil
}

// responsible for get cache data
func (result Result) GetCache(cacheData cache.Cache) (swdata.Cache, error) {
	cache, err := cacheData.GetCache() // call clear cache methode

	if err != nil {
		return swdata.Cache{}, err
	}

	return cache, nil
}
