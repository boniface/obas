package domain

type LocationType struct {
	LocationTypeId string `json:"locationTypeId"`
	Name           string `json:"Name"`
	Code           string `json:"Code"`
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
	LocationId   string       `json:"locationId"`
	Name         string       `json:"Name"`
	LocationType LocationType `json:"locationType"`
	Latitude     string       `json:"Latitude"`
	Longitude    string       `json:"Longitude"`
	ParentId     string       `json:"parentId"`
	//Children     []Location   `json:"Children"`
}
