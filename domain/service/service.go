package service

// package responsible for manage the business rule

import (
	"github.com/SousaGLucas/swsearch/domain/dataprovider/cache" // package responsible for save cache data
	"github.com/SousaGLucas/swsearch/domain/dataprovider/swapi" // package responsible for provide swapi data
	"github.com/SousaGLucas/swsearch/domain/entities/swdata"    // package responsible for mirror database data
)

type Service interface {
	GetData(searchTerm string) (swdata.Data, error) // resposible for get data that each route in the swapi
	ClearCache()                                    // responsible for call clear cache methode
}

// mirror that database data
type Result swdata.Data

// resposible for get data that each route in the swapi
func (result Result) GetData(searchTerm string) (swdata.Data, error) {
	var emptyData swdata.Data // empty variable to return in the error cases

	var cacheData cache.Cache
	cacheData = cache.Data{}

	// check if the term already serached
	if searched := cacheData.CheckTerm(searchTerm); searched != nil {
		return emptyData, searched
	}

	var response swapi.Swapi
	response = swapi.Result{}

	// get response data for the database
	data, err := response.Search(searchTerm)

	if err != nil {
		return emptyData, err
	}

	// adding term in the cache
	cacheData.Push(searchTerm)

	return data, nil
}

// responsible for call clear cache methode
func (result Result) ClearCache() {
	var cacheData cache.Cache
	cacheData = cache.Data{}

	cacheData.Clear() // call clear cache methode
	return
}
