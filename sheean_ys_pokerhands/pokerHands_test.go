package sheenan_ys_pokerhands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func HighCardHandSorted() string{
 return "2C 3D 4D 5C 7C"
}
func LowCardHandSorted() string{
 return "2C 3H 4D 5D 6S"
}
func TestNoGames(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{}))
}

func TestPlayer1WinsWithHighCard(t *testing.T) {
	assert.Equal(t, 1, pokerHands([]string{HighCardHandSorted() + " " + LowCardHandSorted()}))
}
func TestPlayer2WinsWithHighCard(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"2C 3D 4D 5C 6C 2C 3H 4D 5D 7S"}))
}
func TestOneGameUnsorted(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"7C 3D 4D 5C 2C 2C 6H 4D 5D 9S"}))
	assert.Equal(t, 1, pokerHands([]string{"9C 3D 4D 5C 2C 2C 6H 4D 5D 7S"}))
}

func TestOneGameSameRank(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"9C 3D 4D 5C 2C 2C 6H 4D 5D 9S"}))
	assert.Equal(t, 1, pokerHands([]string{"9S 3D 4D 5C 2C 2C 6H 4D 5D 9C"}))
}
func TestOneGameWithCourtCards(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 1, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C"}))
}

func TestOneGamme(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"2C 3D 4D 5C 6C 2C 3H 4D 5D 7S"}))
	assert.Equal(t, 1, pokerHands([]string{"2C 3D 4D 5C 7C 2C 3H 4D 5D 6S"}))
}

func TestTwoGames(t *testing.T) {
	assert.Equal(t, 0, pokerHands([]string{"9C JD 4D 5C 2C KC 6H 4D 5D 8S", "9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 1, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C", "9C JD 4D 5C 2C KC 6H 4D 5D 8S"}))
	assert.Equal(t, 2, pokerHands([]string{"3S AD 4D 5C 2C 2C 6H QD 5D 9C", "3S AD 4D 5C 2C 2C 6H QD 5D 9C"}))
}
