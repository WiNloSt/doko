package pokerhand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const p1WinByHighCard = "2C 3S 4C 5H 8S 2D 3S 4D 5S 7C"
const p2WinByHighCard = "2C 3S 4C 5H 7S 2D 3S 4D 5S 8C"

func TestPokerHand(t *testing.T) {
	t.Run("it should win 0 with no game", func(t *testing.T) {
		assert.Equal(t, 0.0, runGames(""))
	})

	t.Run("it should win 1 with high A", func(t *testing.T) {
		assert.Equal(t, 1.0, runGames(p1WinByHighCard))
	})

	t.Run("it should win 0 with P2 high A", func(t *testing.T) {
		assert.Equal(t, 0.0, runGames(p2WinByHighCard))
	})

	t.Run("player1 win twice", func(t *testing.T) {
		games := p1WinByHighCard + "\n" + p1WinByHighCard
		assert.Equal(t, 1.0, runGames(games))
	})

	t.Run("player1 win once player2 win once", func(t *testing.T) {
		games := p1WinByHighCard + "\n" + p2WinByHighCard
		assert.Equal(t, 0.5, runGames(games))
	})
}

func TestPlayer1Win(t *testing.T) {
	t.Run("player 1 wins by 9 high", func(t *testing.T) {
		assert.Equal(t, true, player1Win("2C 3S 4C 5H 9S 2D 3S 4D 5S 7C"))
	})
}
