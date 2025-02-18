package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0 with no game", func(t *testing.T) {
		assert.Equal(t, 0.0, runGames(""))
	})

	t.Run("it should win 1 with high 8", func(t *testing.T) {
		assert.Equal(t, 1.0, runGames(highestCardWith("8S")+" "+highestCardWith("7C")))
	})

	t.Run("it should win 0 with P2 high 8", func(t *testing.T) {
		assert.Equal(t, 0.0, runGames(highestCardWith("7C")+" "+highestCardWith("8S")))
	})

	t.Run("player1 win twice", func(t *testing.T) {
		games := highestCardWith("8S") + " " + highestCardWith("7C") + "\n" + highestCardWith("8S") + " " + highestCardWith("7C")
		assert.Equal(t, 1.0, runGames(games))
	})

	t.Run("player1 win once player2 win once", func(t *testing.T) {
		games := highestCardWith("8S") + " " + highestCardWith("7C") + "\n" + highestCardWith("7C") + " " + highestCardWith("8S")
		assert.Equal(t, 0.5, runGames(games))
	})
}

func TestPlayer1WinByHighCard(t *testing.T) {
	t.Run("player 1 wins by 9 high", func(t *testing.T) {
		assert.Equal(t, true, player1Win(highestCardWith("9H")+" "+highestCardWith("7C")))
	})

	t.Run("player 1 wins by 10 high", func(t *testing.T) {
		assert.Equal(t, true, player1Win(highestCardWith("TH")+" "+highestCardWith("7C")))
	})

	t.Run("player 1 wins by A in the first order", func(t *testing.T) {
		assert.Equal(t, true, player1Win("AH 2C 3S 4C 5H"+" "+highestCardWith("7C")))
	})

	t.Run("player 1 wins by K in the second order", func(t *testing.T) {
		assert.Equal(t, true, player1Win("2C KH 3S 4C 5H"+" "+highestCardWith("7C")))
	})

	t.Run("player 1 wins by Q in the third order", func(t *testing.T) {
		assert.Equal(t, true, player1Win("2C 3S QH 4C 5H"+" "+highestCardWith("7C")))
	})
}

func highestCardWith(card string) string {
	return "2C 3S 4C 5H " + card
}
