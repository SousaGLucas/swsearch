package swapi

// package responsible for provide swapi data
// responsible for assembling the url for data request to SWAPI

import (
	"errors"
)

// responsible for assembling the url for data request to SWAPI
func request(searchTerm, route string) (map[string]interface{}, error) {
	var emptyData map[string]interface{} // empty variable to return in the error cases

	// getData is being imported from ./swapi.go
	// resp, err := getData("https://swapi.dev/api/" + route + "/?search=" + searchTerm) // call request swapi data function
	resp, err := getData("https://swapi.py4e.com/api/" + route + "/?search=" + searchTerm) // call request swapi data function

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	return resp, nil
}
