package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Input1_say1(t *testing.T){
	answer := FizzBuzz(1)
	expected := "1"
	assert.Equal(t, expected, answer, "ต้องการ 1 ได้ %v", answer)
}