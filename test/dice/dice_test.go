package dice

import (
	"log"
	"os"
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

func TestExternalFile(t *testing.T) {
	b1, err := os.ReadFile("./test1.yaml")
	assert.NoError(t, err)
	log.Println(string(b1))

	b2, err := os.ReadFile("./testdata/test2.yaml")
	assert.NoError(t, err)
	log.Println(string(b2))
}
