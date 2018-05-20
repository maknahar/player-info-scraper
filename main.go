package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"encoding/json"
	"github.com/maknahar/player-info-scraper/player"
	"github.com/maknahar/player-info-scraper/team"
	"io/ioutil"
)

type Teams struct {
	Names []string `json:"names"`
}

func main() {
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	if limit == 0 {
		limit = 1000
	}

	data, err := ioutil.ReadFile("teams.json")
	if err != nil {
		log.Fatalf("Could not read the team.json file. Error: %v", err)
	}

	input := new(Teams)
	err = json.Unmarshal(data, input)
	if err != nil {
		log.Fatalf("Could not decode the team.json file. Error: %v", err)
	}

	teams := team.Scrape(input.Names, limit)
	playersWithTeam, playerWithTeamLookup := make([]player.WithTeam, 0), make(map[string]int)
	for teamName, t := range teams {
		if t == nil {
			log.Println("Unable to find info for team", teamName)
			continue
		}
		for _, p := range t.Players {
			if index, ok := playerWithTeamLookup[p.ID]; ok {
				playersWithTeam[index].Teams = append(playersWithTeam[index].Teams, t.Name)
				continue
			}
			playersWithTeam = append(playersWithTeam, player.WithTeam{Player: p, Teams: []string{t.Name}})
			playerWithTeamLookup[p.ID] = len(playerWithTeamLookup)

		}
	}

	sort.Sort(player.ByName(playersWithTeam))
	for _, playerWithTeam := range playersWithTeam {
		fmt.Println(playerWithTeam)
	}
}
