package domain

type AddressType struct {
	AddressTypeID string `json:"addressTypeID"`
	AddressName   string `json:"addressName"`
}
type ContactType struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type Address struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}
