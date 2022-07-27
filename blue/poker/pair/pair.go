package pair

import "strings"

func IsPair(game string) bool {
	cards := strings.Split(game, " ")
	p1 := cards[0:5]
	// p2 := cards[5:]

	// for i := 0; i < len(p2); i++ {
	// 	for j := i + 1; j < len(p2); j++ {
	// 		returnValue := isMatchCard(p2, i, j)
	// 		if returnValue {
	// 			return returnValue
	// 		}

	// 	}
	// }

	for i := 0; i < len(p1); i++ {
		for j := i + 1; j < len(p1); j++ {
			returnValue := isMatchCard(p1, i, j)
			if returnValue {
				return returnValue
			}

		}
	}

	return false
}

func isMatchCard(p1 []string, i int, j int) bool {
	if p1[i][0] == p1[j][0] {
		return true
	}
	return false
}

// func getMatchCardScore(p1 []string, i int, j int) int {
// 	if p1[i][0] == p1[j][0] {
// 		return p1[i][0]
// 	}
// 	return 0
// }
