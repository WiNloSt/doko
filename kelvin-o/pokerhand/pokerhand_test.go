package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const p2WinByHighCard = "2C 3S 4C 5H 7S 2D 3S 4D 5S 8C"

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0 with no game", func(t *testing.T) {
		assert.Equal(t, 0.0, runGames(""))
	})

	t.Run("it should win 1 with high 8", func(t *testing.T) {
		assert.Equal(t, 1.0, runGames(winsByHighestCardWith("8s")))
	})

	t.Run("it should win 0 with P2 high 8", func(t *testing.T) {
		assert.Equal(t, 0.0, runGames(p2WinByHighCard))
	})

	t.Run("player1 win twice", func(t *testing.T) {
		games := winsByHighestCardWith("8s") + "\n" + winsByHighestCardWith("8s")
		assert.Equal(t, 1.0, runGames(games))
	})

	t.Run("player1 win once player2 win once", func(t *testing.T) {
		games := winsByHighestCardWith("8s") + "\n" + p2WinByHighCard
		assert.Equal(t, 0.5, runGames(games))
	})
}

func TestPlayer1Win(t *testing.T) {
	t.Run("player 1 wins by 9 high", func(t *testing.T) {
		assert.Equal(t, true, player1Win(winsByHighestCardWith("9H")))
	})
}

func winsByHighestCardWith(card string) string {
	return "2C 3S 4C 5H " + card + " 2D 3S 4D 5S 7C"
}
