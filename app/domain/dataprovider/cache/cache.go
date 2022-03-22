package cache

// package responsible for manage cache data

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/SousaGLucas/swsearch/app/domain/entities/swdata" // package responsible for manage the business rule
	"github.com/SousaGLucas/swsearch/log"                        // package responsible for logging system errors
)

type Cache interface {
	CheckTerm(searchTerm string) error // responsible for check if the searchTerm has already been searched
	Push(searchTerm string) error      // responsible for push searcTerm in the cache file
	GetCache() (swdata.Cache, error)   // responsibel for get cache data in the cache file
	Clear() error                      // responsible for clear cache file
}

type Data swdata.Cache

// responsible for check if the searchTerm has already been searched
func (data Data) CheckTerm(searchTerm string) error {
	searched := false

	cache, err := readCache()

	if err != nil {
		return err
	}

	for _, term := range cache {
		if term == searchTerm {
			searched = true
		}
	}

	// checking if searchTerm has already been searched
	if searched {
		return errors.New("'" + searchTerm + "'" + " already searched")
	}

	return nil
}

// responsible for push searcTerm in the cache file
func (data Data) Push(searchTerm string) error {
	cache, err := readCache()

	if err != nil {
		return err
	}

	cache = append(cache, searchTerm)

	writeErr := writeCache(cache)

	if writeErr != nil {
		return err
	}

	return nil
}

// responsibel for get cache data in the cache file
func (data Data) GetCache() (swdata.Cache, error) {
	cache, err := readCache()

	if err != nil {
		return cache, err
	}

	return cache, nil
}

// responsible for clear cache file
func (data Data) Clear() error {
	cache := swdata.Cache{}

	err := writeCache(cache)

	if err != nil {
		return err
	}

	return nil
}

// responsible for read data in the cache file
func readCache() (swdata.Cache, error) {
	var cache swdata.Cache

	readCache, err := os.Open("cache.txt")

	if err != nil {
		log.SetLog(err)
		return cache, errors.New("cache read error")
	}

	defer readCache.Close()

	cacheScanner := bufio.NewScanner(readCache)
	cacheScanner.Split(bufio.ScanLines)

	for cacheScanner.Scan() {
		cache = append(cache, cacheScanner.Text())
	}

	if cacheScanner.Err() != nil {
		log.SetLog(err)
		return cache, errors.New("cache read error")
	}

	return cache, nil
}

// responsible for write data in the cache file
func writeCache(cache swdata.Cache) error {
	cacheFile, err := os.Create("cache.txt")

	if err != nil {
		log.SetLog(err)
		return errors.New("cache create error")
	}

	defer cacheFile.Close()

	cacheWriter := bufio.NewWriter(cacheFile)

	for _, term := range cache {
		fmt.Fprintln(cacheWriter, term)
	}

	if cacheWriter.Flush() != nil {
		log.SetLog(err)
		return errors.New("cache create error")
	}

	return nil
}
