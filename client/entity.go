package client

type ErrorResponse struct {
	Error   int
	Message string
}

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

type Standing struct {
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

type Event struct {
	MatchId                    string       `json:"match_id"`
	CountryId                  string       `json:"country_id"`
	CountryName                string       `json:"country_name"`
	LeagueId                   string       `json:"league_id"`
	LeagueName                 string       `json:"league_name"`
	MatchDate                  string       `json:"match_date"`
	MatchStatus                string       `json:"match_status"`
	MatchTime                  string       `json:"match_time"`
	MatchHometeamName          string       `json:"match_hometeam_name"`
	MatchHometeamScore         string       `json:"match_hometeam_score"`
	MatchAwayteamName          string       `json:"match_awayteam_name"`
	MatchAwayteamScore         string       `json:"match_awayteam_score"`
	MatchHometeamHalftimeScore string       `json:"match_hometeam_halftime_score"`
	MatchAwayteamHalftimeScore string       `json:"match_awayteam_halftime_score"`
	MatchHometeamExtraScore    string       `json:"match_hometeam_extra_score"`
	MatchAwayteamExtraScore    string       `json:"match_awayteam_extra_score"`
	MatchHometeamPenaltyScore  string       `json:"match_hometeam_penalty_score"`
	MatchAwayteamPenaltyScore  string       `json:"match_awayteam_penalty_score"`
	MatchHometeamSystem        string       `json:"match_hometeam_system"`
	MatchAwayteamSystem        string       `json:"match_awayteam_system"`
	MatchLive                  string       `json:"match_live"`
	Goalscorer                 []Goalscorer `json:"goalscorer"`
	Cards                      []Card       `json:"cards"`
	Lineup                     struct {
		Home struct {
			StartingLineups []Lineup         `json:"starting_lineups"`
			Substitutes     []Lineup         `json:"substitutes"`
			Coach           []Lineup         `json:"coach"`
			Substitutions   []LineupWithTime `json:"substitutions"`
		}
		Away struct {
			StartingLineups []Lineup         `json:"starting_lineups"`
			Substitutes     []Lineup         `json:"substitutes"`
			Coach           []Lineup         `json:"coach"`
			Substitutions   []LineupWithTime `json:"substitutions"`
		}
	}
	Statistics []Statistic `json:"statistics"`
}

type Goalscorer struct {
	Time       string `json:"time"`
	HomeScorer string `json:"home_scorer"`
	Score      string `json:"score"`
	AwayScorer string `json:"away_scorer"`
}

type Card struct {
	Time      string `json:"time"`
	HomeFault string `json:"home_fault"`
	Card      string `json:"card"`
	AwayFault string `json:"away_fault"`
}

type Lineup struct {
	LineupPlayer   string `json:"lineup_player"`
	LineupNumber   string `json:"lineup_number"`
	LineupPosition string `json:"lineup_position"`
}

type LineupWithTime struct {
	Lineup
	LineupTime string `json:"lineup_time"`
}

type Statistic struct {
	Type string
	Home string
	Away string
}
