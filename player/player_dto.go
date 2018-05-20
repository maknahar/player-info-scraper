package player

import (
	"fmt"
	"sort"
	"strings"
)

type Player struct {
	Age          string `json:"age"`
	BirthDate    string `json:"birthDate"`
	Country      string `json:"country"`
	FirstName    string `json:"firstName"`
	Height       int    `json:"height"`
	ID           string `json:"id"`
	LastName     string `json:"lastName"`
	Name         string `json:"name"`
	Number       int    `json:"number"`
	Position     string `json:"position"`
	ThumbnailSrc string `json:"thumbnailSrc"`
	Weight       int    `json:"weight"`
}

type WithTeam struct {
	Player
	Teams []string
}

func (p WithTeam) String() string {
	sort.Strings(p.Teams)
	return fmt.Sprintf("%s; %s; %v", p.Name, p.Age, strings.TrimSpace(strings.Join(p.Teams, ", ")))
}
