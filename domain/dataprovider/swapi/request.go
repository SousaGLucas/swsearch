package swapi

// package responsible for provide swapi data
// responsible for get data in the SWAPI

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/SousaGLucas/swsearch/log" // package responsible for manage error log
)

// responsible for get data in the SWAPI
func getData(url string) (map[string]interface{}, error) {
	var emptyData map[string]interface{} // empty variable to return in the error cases

	resp, err := http.Get(url) // get swapi data

	if err != nil {
		log.SetLog(err)
		return emptyData, errors.New("api request error")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.SetLog(err)
		return emptyData, errors.New("error reading data")
	}

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		log.SetLog(err)
		return emptyData, errors.New("error reading data")
	}

	return data, nil
}
