package swapi

// responsible for order and organize SWAPI data, using the search term

import (
	"errors"
)

type Swapi interface {
	Search(searchTerm string) (Result, error)
}

type Result struct {
	People, Planets, Films, Species, Vehicles, Starships map[string]interface{}
}

func (result Result) Search(searchTerm string) (Result, error) {
	// Request is being imported from ./router.go
	people, err := request(searchTerm, "people")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.People = people

	// Request is being imported from ./router.go
	planets, err := request(searchTerm, "planets")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Planets = planets

	// Request is being imported from ./router.go
	films, err := request(searchTerm, "films")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Films = films

	// Request is being imported from ./router.go
	species, err := request(searchTerm, "species")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Species = species

	// Request is being imported from ./router.go
	vehicles, err := request(searchTerm, "vehicles")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Vehicles = vehicles

	// Request is being imported from ./router.go
	starships, err := request(searchTerm, "starships")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Starships = starships

	return result, nil
}
