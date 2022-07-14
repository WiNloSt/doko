package meennew

import "strings"

type Round struct {
	value string
}

func (r Round) isPlayer1Winner() bool {
	sortRank := []string{"A", "K", "Q", "J", "T", "9", "8", "7"}
	for _, rank := range sortRank {
		playerOne := strings.Count(r.value[0:14], rank)
		if playerOne > 1 {
			return true
		}
		count := strings.Count(r.value, rank)
		if count > 1 {
			return false
		}
		founded := strings.Index(r.value, rank)
		if founded < 0 {
			continue
		}
		return foundAtPlayer1(founded)
	}
	return false
}

func foundAtPlayer1(founded int) bool {
	return founded < 14
}
