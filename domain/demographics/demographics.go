package domain

type Role struct {
	Id       string `json:"id"`
	RoleName string `json:"roleName"`
}

type Race struct {
	RaceId   string `json:"raceId"`
	RaceName string `json:"raceName"`
}

type Gender struct {
	GenderId   string `json:"genderId"`
	GenderName string `json:"genderName"`
}

type Title struct {
	TitleId   string `json:"titleId"`
	TitleName string `json:"titlename"`
}

type Province struct {
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
}

type District struct {
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
}

type Town struct {
	TownCode string `json:"townCode"`
	TownName string `json:"townName"`
}

type ProvinceDistrict struct {
	ProvinceCode string `json:"provinceCode"`
	DistrictCode string `json:"districtCode"`
}

type DistrictTown struct {
	DistrictCode string `json:"districtCode"`
	TownCode string `json:"townCode"`
}
