package swdata

// package responsible for manage the business rule

type Data struct {
	People    []People    `json:"people"`
	Planets   []Planets   `json:"planets"`
	Films     []Films     `json:"films"`
	Species   []Species   `json:"species"`
	Vehicles  []Vehicles  `json:"vehicles"`
	Starships []Starships `json:"starships"`
}

type People struct {
	Name      interface{} `json:"name"`
	Height    interface{} `json:"height"`
	Mass      interface{} `json:"mass"`
	HairColor interface{} `json:"hair_color"`
	SkinColor interface{} `json:"skin_color"`
	EyeColor  interface{} `json:"eye_color"`
	BirthYear interface{} `json:"birth_year"`
	Gender    interface{} `json:"gender"`
}

type Planets struct {
	Name           interface{} `json:"name"`
	RotationPeriod interface{} `json:"rotate_period"`
	OrbitalPeriod  interface{} `json:"orbital_period"`
	Diameter       interface{} `json:"diameter"`
	Climate        interface{} `json:"climate"`
	Gravity        interface{} `json:"gravity"`
	Terrain        interface{} `json:"terrain"`
	SurfaceWater   interface{} `json:"surface_water"`
	Population     interface{} `json:"population"`
}

type Films struct {
	Title        interface{} `json:"title"`
	EpisodeID    interface{} `json:"episode_id"`
	OpeningCrawl interface{} `json:"opening_crawl"`
	Director     interface{} `json:"director"`
	Producer     interface{} `json:"producer"`
	ReleaseDate  interface{} `json:"release_date"`
}

type Species struct {
	Name            interface{} `json:"name"`
	Classification  interface{} `json:"classification"`
	Designation     interface{} `json:"designation"`
	AverageHeight   interface{} `json:"average_height"`
	SkinColors      interface{} `json:"skin_colors"`
	HairColors      interface{} `json:"hair_colors"`
	EyeColors       interface{} `json:"eye_colors"`
	AverageLifespan interface{} `json:"average_lifespan"`
	Language        interface{} `json:"language"`
}

type Vehicles struct {
	Name                 interface{} `json:"name"`
	Model                interface{} `json:"model"`
	Manufacturer         interface{} `json:"manufacturer"`
	CostInCredits        interface{} `json:"cost_in_credits"`
	Length               interface{} `json:"lenght"`
	MaxAtmospheringSpeed interface{} `json:"max_atmosphering_speed"`
	Crew                 interface{} `json:"crew"`
	Passengers           interface{} `json:"passengers"`
	CargoCapacity        interface{} `json:"cargo_capacity"`
	Consumables          interface{} `json:"consumables"`
	VehicleClass         interface{} `json:"vehicle_class"`
}

type Starships struct {
	Name                 interface{} `json:"name"`
	Model                interface{} `json:"model"`
	Manufacturer         interface{} `json:"manufacturer"`
	CostInCredits        interface{} `json:"cost_in_credits"`
	Length               interface{} `json:"length"`
	MaxAtmospheringSpeed interface{} `json:"max_atmosphering_speed"`
	Crew                 interface{} `json:"crew"`
	Passengers           interface{} `json:"passengers"`
	CargoCapacity        interface{} `json:"cargo_capacity"`
	Consumables          interface{} `json:"consumables"`
	HyperdriveRating     interface{} `json:"hyperdrive_rating"`
	MGLT                 interface{} `json:"MGLT"`
	StarshipClass        interface{} `json:"starship_class"`
}

type Cache []string
