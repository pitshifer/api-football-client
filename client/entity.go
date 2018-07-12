package client

type Country struct {
	Id   string `json:"country_id"`
	Name string `json:"country_name"`
}

type League struct {
	CountryId   string `json:"country_id"`
	CountryName string `json:"country_name"`
	Id          string `json:"league_id"`
	Name        string `json:"league_name"`
}
