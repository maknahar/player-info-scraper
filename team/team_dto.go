package team

import "github.com/maknahar/player-info-scraper/player"

// Response contain common fields among HTTP response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// APIResponse Store the content from onefootball team API
type APIResponse struct {
	Response
	Data struct {
		Team `json:"team"`
	} `json:"data"`
}

// Team contain information about a team
type Team struct {
	Colors       Color           `json:"colors"`
	Competitions []Competition   `json:"competitions"`
	ID           int             `json:"id"`
	IsNational   bool            `json:"isNational"`
	LogoUrls     []LogoUrl       `json:"logoUrls"`
	Name         string          `json:"name"`
	Officials    []Official      `json:"officials"`
	OptaID       int             `json:"optaId"`
	Players      []player.Player `json:"players"`
}

type Color struct {
	CrestMainColor string `json:"crestMainColor"`
	MainColor      string `json:"mainColor"`
	ShirtColorAway string `json:"shirtColorAway"`
	ShirtColorHome string `json:"shirtColorHome"`
}

type Competition struct {
	CompetitionID   int    `json:"competitionId"`
	CompetitionName string `json:"competitionName"`
}

type LogoUrl struct {
	Size string `json:"size"`
	URL  string `json:"url"`
}

type Official struct {
	Country     string `json:"country"`
	CountryName string `json:"countryName"`
	FirstName   string `json:"firstName"`
	ID          string `json:"id"`
	LastName    string `json:"lastName"`
	Position    string `json:"position"`
}
