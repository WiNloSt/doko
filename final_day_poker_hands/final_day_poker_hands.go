package final_day_poker_hands

import (
	"strconv"
	"strings"
	"sort"
)

type Cards []string

func pokerhands(games []string) int {
	p1WinCount := 0
	for _, game := range games {
		cards := strings.Split(game, " ")
		if P1WinsCompareHighCard(cards) || P1WinsOnePair(cards) {
			p1WinCount += 1
		}

	}
	return p1WinCount
}

func getFaceValue(face string) int {
	switch face {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	}
	faceValue, _ := strconv.Atoi(face)
	return faceValue
}

func P1WinsOnePair(cards []string) bool {
	return cards[3][0] == cards[4][0]
}

type Hand []string

func (hand1 Hand) isBiggerThan(hand2 Hand)bool{
	for i := 4; i >= 0; i-- {
		if hand1[i][:1] != hand2[i][:1] {
			return hand1[i][:1] > hand2[i][:1]
		}
	}
	return false
}

func P1WinsCompareHighCard(game []string) bool {

	if getFaceValue(game[4][:1]) == 14 {
		return true
	}
	if getFaceValue(game[4][:1]) == 13 {
		return true
	}

	hand1,hand2 := game[:5],game[5:]
	if Hand(hand1).isBiggerThan(Hand(hand2)){
		return true
	}

	return false
}

func (c Cards) sort() Cards {
	sorted := sort.StringSlice(c)
	sorted.Sort()
	return Cards(sorted)
}
