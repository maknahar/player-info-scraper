package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/maknahar/player-info-scraper/player"
	"github.com/maknahar/player-info-scraper/team"
)

func main() {
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	if limit == 0 {
		//log.Println("LIMIT is not set. Defaulting to 1000")
		limit = 1000
	}

	teamNames := []string{"Germany", "England", "France", "Spain", "Manchester Utd", "Arsenal", "Chelsea",
		"Barcelona", "Real Madrid", "FC Bayern Munich"}

	teams := team.Scrape(teamNames, limit)
	playersWithTeam, playerWithTeamLookup := make([]player.WithTeam, 0), make(map[string]int)
	for _, t := range teams {
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
