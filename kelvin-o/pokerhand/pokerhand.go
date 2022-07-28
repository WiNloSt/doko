package pokerhand

import "strings"

func runGames(games string) float64 {
	if len(games) == 0 {
		return 0.0
	}

	winCount := 0.0
	gameCount := 0.0
	for _, game := range strings.Split(games, "\n") {
		gameCount += 1
		if player1Win(game) {
			winCount += 1
		}
	}

	return winCount / gameCount
}

func player1Win(game string) bool {
	p1Hand := Hand{cards: game[0:14]}
	p2Hand := Hand{cards: game[15:]}
	return getRank(p1Hand.getCardRank(4)) > getRank(p2Hand.getCardRank(4))
}

func getRank(cardRank byte) int {
	return strings.Index("23456789TJQKA", string(cardRank))
}

type Hand struct {
	cards string
}

func (hand *Hand) getCardRank(cardIndex int) byte {
	return hand.cards[cardIndex*3]
}
