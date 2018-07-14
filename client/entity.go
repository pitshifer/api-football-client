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

type Standings struct {
	CountryName           string `json:"country_name"`
	LeagueId              string `json:"league_id"`
	LeagueName            string `json:"league_name"`
	TeamName              string `json:"team_name"`
	OverallLeaguePosition string `json:"overall_league_position"`
	OverallLeaguePayed    string `json:"overall_league_payed"`
	OverallLeagueW        string `json:"overall_league_w"`
	OverallLeagueD        string `json:"overall_league_d"`
	OverallLeagueL        string `json:"overall_league_l"`
	OverallLeagueGF       string `json:"overall_league_gf"`
	OverallLeagueGA       string `json:"overall_league_ga"`
	OverallLeaguePTS      string `json:"overall_league_pts"`
	HomeLeaguePosition    string `json:"home_league_position"`
	HomeLeaguePayed       string `json:"home_league_payed"`
	HomeLeagueW           string `json:"home_league_w"`
	HomeLeagueD           string `json:"home_league_d"`
	HomeLeagueL           string `json:"home_league_l"`
	HomeLeagueGF          string `json:"home_league_gf"`
	HomeLeagueGA          string `json:"home_league_ga"`
	HomeLeaguePTS         string `json:"home_league_pts"`
	AwayLeaguePosition    string `json:"away_league_position"`
	AwayLeaguePayed       string `json:"away_league_payed"`
	AwayLeagueW           string `json:"away_league_w"`
	AwayLeagueD           string `json:"away_league_d"`
	AwayLeagueL           string `json:"away_league_l"`
	AwayLeagueGF          string `json:"away_league_gf"`
	AwayLeagueGA          string `json:"away_league_ga"`
	AwayLeaguePTS         string `json:"away_league_pts"`
}
