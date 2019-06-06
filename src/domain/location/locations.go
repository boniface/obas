package location

type LocationType struct {
	Id   string
	Name string
	Code string
}

type AddressType struct {
	Id   string
	Name string
}
type ContactType struct {
	Id   string
	Name string
}

type Location struct {
	Id           string
	Name         string
	LocationType LocationType
	Latitude     string
	Longitude    string
	Larent       Location
	Children     []Location
}
