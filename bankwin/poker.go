package meennew

func PokerHand(records []string) int {
	countP1Winner := 0
	if len(records) == 0 {
		return countP1Winner
	}
	for _, value := range records {
		round := Round{value}
		if round.isPlayer1Winner() {
			countP1Winner += 1
		}
	}
	return countP1Winner
}