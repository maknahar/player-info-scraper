package team

import (
	"testing"

	utils "github.com/maknahar/go-utils"
)

func TestScrape(t *testing.T) {
	teams := Scrape([]string{"Apoel FC"}, 1)
	utils.ShouldBeEqual(t, "Apoel FC", teams["Apoel FC"].Name)
	utils.ShouldBeEqual(t, 1, len(teams))

	teams = Scrape([]string{"Apoel FC"}, 100)
	utils.ShouldBeEqual(t, "Apoel FC", teams["Apoel FC"].Name)
	utils.ShouldBeEqual(t, 1, len(teams))

	teams = Scrape([]string{"Apoel FC", "Manchester United"}, 2)
	utils.ShouldBeEqual(t, "Apoel FC", teams["Apoel FC"].Name)
	utils.ShouldBeEqual(t, 2, len(teams))
	if teams["Manchester United"] != nil {
		t.Fatalf("Expected %v. Got %v", nil, teams["Manchester United"])
	}

	teams = Scrape([]string{"Germany", "England", "France", "Spain", "Manchester Utd", "Arsenal", "Chelsea",
		"Barcelona", "Real Madrid", "FC Bayern Munich"}, 100)
	utils.ShouldBeEqual(t, 10, len(teams))
	for k, v := range teams {
		if v == nil {
			t.Fatalf("Cound not find the team %s", k)
		}
	}

	teams = Scrape([]string{"Germany", "England", "France", "Spain", "Manchester Utd", "Arsenal", "Chelsea",
		"Barcelona", "Real Madrid", "FC Bayern Munich"}, 1000)
	utils.ShouldBeEqual(t, 10, len(teams))
	for k, v := range teams {
		if v == nil {
			t.Fatalf("Cound not find the team %s", k)
		}
	}
}
