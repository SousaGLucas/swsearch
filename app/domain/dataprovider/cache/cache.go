package cache

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/SousaGLucas/swsearch/app/domain/entities/swdata"
	"github.com/SousaGLucas/swsearch/log"
)

type Cache interface {
	CheckTerm(searchTerm string) error
	Push(searchTerm string) error
	GetCache() (swdata.Cache, error)
	Clear() error
}

type Data swdata.Cache

var cache swdata.Cache

func (data Data) CheckTerm(searchTerm string) error {
	searched := false

	err := readCache()

	if err != nil {
		return err
	}

	for _, term := range cache {
		if term == searchTerm {
			searched = true
		}
	}

	if searched {
		return errors.New(searchTerm + " already searched")
	}

	return nil
}

func (data Data) Push(searchTerm string) error {
	cache = append(cache, searchTerm)

	err := writeCache()

	if err != nil {
		return err
	}

	return nil
}

func (data Data) GetCache() (swdata.Cache, error) {
	emptyCache := swdata.Cache{}

	err := readCache()

	if err != nil {
		return emptyCache, err
	}

	return cache, nil
}

func (data Data) Clear() error {
	cache = []string{}

	err := writeCache()

	if err != nil {
		return err
	}

	return nil
}

func readCache() error {
	readCache, err := os.Open("cache.txt")

	if err != nil {
		log.SetLog(err)
		return errors.New("cache read error")
	}

	defer readCache.Close()

	cacheScanner := bufio.NewScanner(readCache)
	cacheScanner.Split(bufio.ScanLines)

	for cacheScanner.Scan() {
		cache = append(cache, cacheScanner.Text())
	}

	if cacheScanner.Err() != nil {
		log.SetLog(err)
		return errors.New("cache read error")
	}

	return nil
}

func writeCache() error {
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
