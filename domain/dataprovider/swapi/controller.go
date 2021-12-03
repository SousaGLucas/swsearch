package swapi

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
	people, err := Request(searchTerm, "people")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.People = people

	// Request is being imported from ./router.go
	planets, err := Request(searchTerm, "planets")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Planets = planets

	// Request is being imported from ./router.go
	films, err := Request(searchTerm, "films")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Films = films

	// Request is being imported from ./router.go
	species, err := Request(searchTerm, "species")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Species = species

	// Request is being imported from ./router.go
	vehicles, err := Request(searchTerm, "vehicles")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Vehicles = vehicles

	// Request is being imported from ./router.go
	starships, err := Request(searchTerm, "starships")

	if err != nil {
		var emptyResult Result
		return emptyResult, errors.New("api request error")
	}

	result.Starships = starships

	return result, nil
}
