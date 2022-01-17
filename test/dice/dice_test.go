package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		out := PlayDice()
		t.Log(out)
		assert.GreaterOrEqual(t, out, 1)
		assert.LessOrEqual(t, out, 6)
	}
}
