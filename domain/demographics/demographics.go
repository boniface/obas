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
