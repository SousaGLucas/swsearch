package main

import (
	"encoding/json"
	"fmt"

	"github.com/SousaGLucas/swsearch/domain/dataprovider/swapi"
)

func main() {
	var search swapi.Swapi

	search = &swapi.Result{}

	searchTerm := "sky"

	result, err := search.Search(searchTerm)

	if err != nil {
		return
	}

	people, _ := json.Marshal(result.People["results"])
	planets, _ := json.Marshal(result.Planets["results"])
	films, _ := json.Marshal(result.Films["results"])
	species, _ := json.Marshal(result.Species["results"])
	vehicles, _ := json.Marshal(result.Vehicles["results"])
	starships, _ := json.Marshal(result.Starships["results"])

	fmt.Printf("\nResults for '%v':\n\n", searchTerm)

	fmt.Printf("People:\n\n %v\n\n", string(people))
	fmt.Printf("Planets:\n\n %v\n\n", string(planets))
	fmt.Printf("Films:\n\n %v\n\n", string(films))
	fmt.Printf("Species:\n\n %v\n\n", string(species))
	fmt.Printf("Vehicles:\n\n %v\n\n", string(vehicles))
	fmt.Printf("Starships:\n\n %v\n\n", string(starships))
}
