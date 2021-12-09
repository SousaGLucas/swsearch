package cache

import "errors"

type Cache interface {
	CheckTerm(searchTerm string) error
	Push(searchTerm string)
	Clear()
}

type Data map[string]int

var cache Data = make(Data)

func (data Data) CheckTerm(searchTerm string) error {
	if _, ok := cache[searchTerm]; ok {
		return errors.New(searchTerm + " already searched")
	}

	return nil
}

func (data Data) Push(searchTerm string) {
	length := 0
	for range cache {
		length += 1
	}

	cache[searchTerm] = length + 1
	return
}

func GetCache() Data {
	return cache
}

func (data Data) Clear() {
	for i := range cache {
		delete(cache, i)
	}
}
