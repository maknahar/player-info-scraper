package team

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var teamAPIEndPoint string

func init() {
	teamAPIEndPoint = os.Getenv("TEAM_API_ENDPOINT")
	if teamAPIEndPoint == "" {
		teamAPIEndPoint = "https://vintagemonster.onefootball.com/api/teams/en/%d.json"
	}
}

// Scrape scrape the team API and find the teams
// limit is an integer which can limit the api calls to find the teams
func Scrape(teamNames []string, limit int) (teams map[string]*Team) {
	//t := time.Now()
	//defer log.Println("Scraping finished in ", time.Since(t))

	// Create and initialize teams
	teams = make(map[string]*Team)
	for _, teamName := range teamNames {
		teams[teamName] = nil
	}

	availableTeamCount := 0
	wg := sync.WaitGroup{}
	for i := 1; availableTeamCount != len(teams) && i <= limit; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//log.Println("Getting team", i)
			team, err := Get(i)
			if err != nil {
				log.Printf("Error in getting team info with ID %d . %v", i, err)
				return
			}

			if _, ok := teams[team.Name]; ok {
				teams[team.Name] = team
				availableTeamCount++
			}
		}(i)
		// Get team details in batches to avoid overwhelming the server
		if i%50 == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
	return teams
}

// Get retrieve the team info
func Get(teamID int) (*Team, error) {
	resp, err := http.Get(fmt.Sprintf(teamAPIEndPoint, teamID))
	if err != nil {
		return nil, fmt.Errorf("error in getting the team info: %v+", err)
	}
	defer resp.Body.Close()

	team := new(APIResponse)
	err = json.NewDecoder(resp.Body).Decode(team)
	if err != nil {
		return nil, fmt.Errorf("error in decoding the team info: %v+", err)
	}

	if team.Response.Code > http.StatusOK {
		return nil, fmt.Errorf("got invalid response code: %d", team.Response.Code)
	}

	return &team.Data.Team, nil
}
