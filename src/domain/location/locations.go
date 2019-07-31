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
	LocationId     string `json:"locationId"`
	Name           string `json:"name"`
	LocationTypeId string `json:"locationTypeId"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	ParentId       string `json:"parentId"`
	//Children     []Location   `json:"Children"`
}
