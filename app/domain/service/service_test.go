package service

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/cache"
	"github.com/SousaGLucas/swsearch/app/domain/dataprovider/swapi"
	"github.com/SousaGLucas/swsearch/app/domain/entities/swdata"
)

var expectedErr = errors.New("expected error")

type mockCache struct {
	GetCacheFn  func() (swdata.Cache, error)
	ClearFn     func() error
	CheckTermFn func(searchTerm string) error
}

func (m mockCache) CheckTerm(searchTerm string) error {
	return m.CheckTermFn(searchTerm)
}

func (m mockCache) Push(searchTerm string) error {
	return m.CheckTermFn(searchTerm)
}

func (m mockCache) GetCache() (swdata.Cache, error) {
	return m.GetCacheFn()
}

func (m mockCache) Clear() error {
	return m.ClearFn()
}

type mockSwapi struct {
	SearchFn func(searchTerm string) (swdata.Data, error)
}

func (m mockSwapi) Search(searchTerm string) (swdata.Data, error) {
	return m.SearchFn(searchTerm)
}

func TestGetCache(t *testing.T) {
	cacheTest := []string{"Sky", "r2", "Lea"}

	testCases := []struct {
		name   string
		cache  cache.Cache
		err    error
		result swdata.Cache
	}{
		{
			name: "returns slice of result should be success",
			cache: mockCache{
				GetCacheFn: func() (swdata.Cache, error) {
					return swdata.Cache(cacheTest), nil
				},
			},
			result: swdata.Cache(cacheTest),
			err:    nil,
		},
		{
			name: "returns err should return err",
			cache: mockCache{
				GetCacheFn: func() (swdata.Cache, error) {
					return nil, expectedErr
				},
			},
			result: swdata.Cache([]string{}),
			err:    expectedErr,
		},
	}

	response := Result{}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			cache, err := response.GetCache(testCase.cache)

			// wait error
			if testCase.err != nil {
				if testCase.err != err {
					t.Errorf("expected error %v, got %v", testCase.err, err)
				}
				return
			}

			jsonCache, err := json.Marshal(cache)
			if err != nil {
				t.Errorf("json convetion error")
			}

			jsonResult, err := json.Marshal(testCase.result)
			if err != nil {
				t.Errorf("json convetion error")
			}

			if string(jsonCache) != string(jsonResult) {
				t.Errorf("expected %v, return %v", testCase.cache, cache)
			}
		})

	}
}

func TestClearCache(t *testing.T) {
	testCases := []struct {
		name  string
		cache cache.Cache
		err   error
	}{
		{
			name: "returns nil should be success",
			cache: mockCache{
				ClearFn: func() error {
					return nil
				},
			},
			err: nil,
		},
		{
			name: "returns err should return err",
			cache: mockCache{
				ClearFn: func() error {
					return expectedErr
				},
			},
			err: expectedErr,
		},
	}

	response := Result{}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			err := response.ClearCache(testCase.cache)

			if testCase.err != err {
				t.Errorf("expected error %v, got %v", testCase.err, err)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	testCases := []struct {
		name     string
		term     string
		cache    cache.Cache
		response swapi.Swapi
		result   swdata.Data
		err      error
	}{
		{
			name: "returns err of term already searched should return err",
			term: "term",
			cache: mockCache{
				CheckTermFn: func(searchTerm string) error {
					return expectedErr
				},
			},
			response: mockSwapi{
				SearchFn: func(searchTerm string) (swdata.Data, error) {
					return swdata.Data{}, expectedErr
				},
			},
			err: expectedErr,
		},
		{
			name: "returns err should return err",
			term: "term",
			cache: mockCache{
				CheckTermFn: func(searchTerm string) error {
					return expectedErr
				},
			},
			response: mockSwapi{
				SearchFn: func(searchTerm string) (swdata.Data, error) {
					return swdata.Data{}, expectedErr
				},
			},
			err: expectedErr,
		},

		// {
		// 	name: "returns err should return err",
		// 	cache: mockCache{
		// 		ClearFn: func() error {
		// 			return expectedErr
		// 		},
		// 	},
		// 	err: expectedErr,
		// },
	}

	response := Result{}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			_, err := response.GetData(testCase.term, testCase.response, testCase.cache)

			if testCase.err != err {
				t.Errorf("expected error %v, got %v", testCase.err, err)
			}
		})
	}
}
