package swdata

// package responsible for mirror database data

type Data struct {
	People    map[string]interface{}
	Planets   map[string]interface{}
	Films     map[string]interface{}
	Species   map[string]interface{}
	Vehicles  map[string]interface{}
	Starships map[string]interface{}
}

type Cache []string
