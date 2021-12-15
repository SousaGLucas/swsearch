package swapi

// package responsible for provide swapi data
// responsible for order and organize SWAPI data, using the search term

import (
	"errors"

	"github.com/SousaGLucas/swsearch/app/domain/entities/swdata" // package responsible for mirror database data
)

type Swapi interface {
	Search(searchTerm string) (swdata.Data, error) // responsible for call serach swapi data function for each api route
}

// mirror that database data
type Result swdata.Data

// responsible for call serach swapi data function for each api route
func (result Result) Search(searchTerm string) (swdata.Data, error) {
	var emptyData swdata.Data // empty variable to return in the error cases
	var data swdata.Data      // variable for save api data

	// request is being imported from ./router.go
	people, err := request(searchTerm, "people") // get data for the 'people' route

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	data.People = people

	// request is being imported from ./router.go
	planets, err := request(searchTerm, "planets") // get data for the 'planets' route

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	data.Planets = planets

	// request is being imported from ./router.go
	films, err := request(searchTerm, "films") // get data for the 'films' route

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	data.Films = films

	// request is being imported from ./router.go
	species, err := request(searchTerm, "species") // get data for the 'species' route

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	data.Species = species

	// request is being imported from ./router.go
	vehicles, err := request(searchTerm, "vehicles") // get data for the 'vehicles' route

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	data.Vehicles = vehicles

	// request is being imported from ./router.go
	starships, err := request(searchTerm, "starships") // get data for the 'starchips' route

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	data.Starships = starships

	return data, nil
}
