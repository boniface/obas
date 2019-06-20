package domain

type LocationType struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Code string `json:"Code"`
}

type AddressType struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}
type ContactType struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type Location struct {
	Id           string       `json:"Id"`
	Name         string       `json:"Name"`
	LocationType LocationType `json:"locationType"`
	Latitude     string       `json:"Latitude"`
	Longitude    string       `json:"Longitude"`
	Larent       Location     `json:"Larent"`
	Children     []Location   `json:"Children"`
}
