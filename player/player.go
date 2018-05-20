package player

import "sort"

// ByName allow players to be sorted by name
type ByName []WithTeam

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return sort.StringsAreSorted([]string{n[i].Name, n[j].Name}) }
