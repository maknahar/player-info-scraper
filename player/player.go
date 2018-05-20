package player

import (
	"log"
	"sort"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var t = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

// ByName allow players to be sorted by name
type ByName []WithTeam

func (n ByName) Len() int      { return len(n) }
func (n ByName) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool {
	name1, _, err := transform.String(t, n[i].Name)
	if err != nil {
		log.Println("Could not transform", n[i].Name, err)
		name1 = n[i].Name
	}
	name2, _, err := transform.String(t, n[j].Name)
	if err != nil {
		log.Println("Could not transform", n[j].Name, err)
		name2 = n[j].Name
	}

	return sort.StringsAreSorted([]string{name1, name2})
}
