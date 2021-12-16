package swapi

// package responsible for provide swapi data
// responsible for order and organize SWAPI data, using the search term

import "github.com/SousaGLucas/swsearch/app/domain/entities/swdata" // package responsible for mirror database data

type Swapi interface {
	Search(searchTerm string) (swdata.Data, error) // responsible for call serach swapi data function for each api route
}

// mirror that database data
type Result swdata.Data

// responsible for call serach swapi data function for each api route
func (result Result) Search(searchTerm string) (swdata.Data, error) {
	emptyData := swdata.Data{} // empty variable to return in the error cases
	data := swdata.Data{}      // variable for save api data

	// request is being imported from ./router.go
	people, err := request(searchTerm, "people") // get data for the 'people' route

	if err != nil {
		return emptyData, err
	}

	for _, item := range people {
		elem := swdata.People{}

		elem.Name = item["name"]
		elem.Height = item["height"]
		elem.Mass = item["mass"]
		elem.HairColor = item["hair_color"]
		elem.SkinColor = item["skin_color"]
		elem.EyeColor = item["eye_color"]
		elem.BirthYear = item["birth_year"]
		elem.Gender = item["gender"]

		data.People = append(data.People, elem)
	}

	// request is being imported from ./router.go
	planets, err := request(searchTerm, "planets") // get data for the 'planets' route

	if err != nil {
		return emptyData, err
	}

	for _, item := range planets {
		elem := swdata.Planets{}

		elem.Name = item["name"]
		elem.RotationPeriod = item["rotation_period"]
		elem.OrbitalPeriod = item["orbital_period"]
		elem.Diameter = item["diameter"]
		elem.Climate = item["climate"]
		elem.Gravity = item["gravity"]
		elem.Terrain = item["terrain"]
		elem.SurfaceWater = item["surface_water"]
		elem.Population = item["population"]

		data.Planets = append(data.Planets, elem)
	}

	// request is being imported from ./router.go
	films, err := request(searchTerm, "films") // get data for the 'films' route

	if err != nil {
		return emptyData, err
	}

	for _, item := range films {
		elem := swdata.Films{}

		elem.Title = item["title"]
		elem.EpisodeID = item["episode_id"]
		elem.OpeningCrawl = item["opening_crawl"]
		elem.Director = item["director"]
		elem.Producer = item["producer"]
		elem.ReleaseDate = item["release_date"]

		data.Films = append(data.Films, elem)
	}

	// request is being imported from ./router.go
	species, err := request(searchTerm, "species") // get data for the 'species' route

	if err != nil {
		return emptyData, err
	}

	for _, item := range species {
		elem := swdata.Species{}

		elem.Name = item["name"]
		elem.Classification = item["classification"]
		elem.Designation = item["designation"]
		elem.AverageHeight = item["average_height"]
		elem.SkinColors = item["skin_colors"]
		elem.HairColors = item["hair_colors"]
		elem.EyeColors = item["eye_colors"]
		elem.AverageLifespan = item["average_lifespan"]
		elem.Language = item["language"]

		data.Species = append(data.Species, elem)
	}

	// request is being imported from ./router.go
	vehicles, err := request(searchTerm, "vehicles") // get data for the 'vehicles' route

	if err != nil {
		return emptyData, err
	}

	for _, item := range vehicles {
		elem := swdata.Vehicles{}

		elem.Name = item["name"]
		elem.Model = item["model"]
		elem.Manufacturer = item["manufacturer"]
		elem.CostInCredits = item["coin_in_credits"]
		elem.Length = item["length"]
		elem.MaxAtmospheringSpeed = item["max_atmosphering_speed"]
		elem.Crew = item["crew"]
		elem.Passengers = item["passengers"]
		elem.CargoCapacity = item["cargo_capacity"]
		elem.Consumables = item["consumables"]
		elem.VehicleClass = item["vehicle_class"]

		data.Vehicles = append(data.Vehicles, elem)
	}

	// request is being imported from ./router.go
	starships, err := request(searchTerm, "starships") // get data for the 'starchips' route

	if err != nil {
		return emptyData, err
	}

	for _, item := range starships {
		elem := swdata.Starships{}

		elem.Name = item["name"]
		elem.Model = item["model"]
		elem.Manufacturer = item["manufacturer"]
		elem.CostInCredits = item["coin_in_credits"]
		elem.Length = item["length"]
		elem.MaxAtmospheringSpeed = item["max_atmosphering_speed"]
		elem.Crew = item["crew"]
		elem.Passengers = item["passengers"]
		elem.CargoCapacity = item["cargo_capacity"]
		elem.Consumables = item["consumables"]
		elem.HyperdriveRating = item["hyperdrive_rating"]
		elem.MGLT = item["MGLT"]
		elem.StarshipClass = item["starship_class"]

		data.Starships = append(data.Starships, elem)
	}

	return data, nil
}
