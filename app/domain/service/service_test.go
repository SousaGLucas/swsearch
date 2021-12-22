package service

import (
	"encoding/json"
	"testing"

	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/cache"
)

func TestGetCache(t *testing.T) {
	cacheExpected, err := json.Marshal([]string{"Sky", "r2", "Lea"})

	if err != nil {
		return
	}

	testCases := []struct {
		cache string
		err   error
	}{
		{string(cacheExpected), nil},
	}

	response := Result{}
	cacheData := cache.MockData{}

	for _, testCase := range testCases {
		cache, err := response.GetCache(cacheData)

		if err != testCase.err {
			t.Errorf("Expected nil error")
		}

		jsonCache, err := json.Marshal(cache)

		if err != nil {
			t.Errorf("json convetion error")
		}

		if string(jsonCache) != testCase.cache {
			t.Errorf("Expected %v, return %v", testCase.cache, string(jsonCache))
		}
	}
}
