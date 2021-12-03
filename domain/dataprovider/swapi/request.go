package swapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/SousaGLucas/swsearch/log"
)

func GetData(url string) (map[string]interface{}, error) {
	emptyData := make(map[string]interface{})

	resp, err := http.Get(url)

	if err != nil {
		log.SetLog(err)
		return emptyData, errors.New("api request error")
	}

	// defer resp.Body.Close()

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
