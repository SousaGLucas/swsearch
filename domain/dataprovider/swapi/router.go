package swapi

import "errors"

func Request(searchTerm, route string) (map[string]interface{}, error) {
	emptyData := make(map[string]interface{})

	// getData is being imported from ./swapi.go
	resp, err := GetData("https://swapi.dev/api/" + route + "/?search=" + searchTerm)

	if err != nil {
		return emptyData, errors.New("api request error")
	}

	return resp, nil
}
