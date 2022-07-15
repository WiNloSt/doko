package meennew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoRecord(t *testing.T) {
	records := []string{}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestOneRecord(t *testing.T) {
	records := []string{"5S AD TS 3H 2S 4H 2H 5H 6S 9S"}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestTwoRecord(t *testing.T) {
	records := []string{
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestPlayerOneLose(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S KS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestPlayerOneWinInTwoRound(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S QS",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestPlayer2WinOf3Round(t *testing.T) {
	records := []string{
		"5S 6D TS 3H 2S 4H 2H 5H 6S QS",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
		"5S AD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 2)
}

func TestCompareP1WinHighCard(t *testing.T) {
	records := []string{
		"5S KD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP1WinHighCardWithQueen(t *testing.T) {
	records := []string{
		"5S QD TS 3H 2S 4H 2H 5H 6S 9S",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCompareP2WinHighCardWithQueen(t *testing.T) {
	records := []string{
		"5S KS TS 3H 2S 4H 2H 5H 6S AS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 0)
}

func TestCompareP1WinHighCardWithJack(t *testing.T) {
	records := []string{
		"5S JS TS 3H 2S 4H 2H 5H 6S TS",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 1)
}

func TestCountP1(t *testing.T) {
	records := []string{
		"5S 9S TS 3H 2S 4H 2H 5H 6S 7S",
		"5S 9S TS 3H 2S 4H 2H 5H 6S 7S",
		"5S 9S KS 3H 2S 4H 2H 5H 6S 7S",
		"AC JC 5S 3H 2S 4C 2H 5H AH JH",
		"6C 4C 5S 3H 9S 4C 2H 5H 8H AH",
	}
	winRateP1 := PokerHand(records)
	assert.Equal(t, winRateP1, 3)
}