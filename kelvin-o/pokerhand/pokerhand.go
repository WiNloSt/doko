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
	for index := range p1Hand.getCards() {
		if index == 4 {
			continue
		}
		if p1Hand.getCardRank(index) > p1Hand.getCardRank(index+1) {
			p1Hand.moveCard(index, index+1)
		}
	}
	return getRank(p1Hand.getCardRank(4)) > getRank(p2Hand.getCardRank(4))
}

func (hand *Hand) moveCard(fromIndex int, toIndex int) {
	hand.cards += " "
	hand.cards = hand.cards[0*3:fromIndex*3] + hand.cards[(fromIndex+1)*3:(toIndex+1)*3] + hand.cards[fromIndex*3:(fromIndex+1)*3] + hand.cards[(toIndex+1)*3:]
	hand.cards = hand.cards[0:14]
}

func getRank(cardRank byte) int {
	return strings.Index("23456789TJQKA", string(cardRank))
}

type Hand struct {
	cards string
}

func (hand *Hand) getCardRank(cardIndex int) byte {
	return hand.getCards()[cardIndex][0]
}

func (hand *Hand) getCards() []string {
	return strings.Split(hand.cards, " ")
}
