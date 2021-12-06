package swapi

// responsible for assembling the url for data request to SWAPI

import "errors"

func request(searchTerm, route string) (map[string]interface{}, error) {
	emptyData := make(map[string]interface{})

	// getData is being imported from ./swapi.go
	resp, err := getData("https://swapi.dev/api/" + route + "/?search=" + searchTerm)

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	return resp, nil
}
