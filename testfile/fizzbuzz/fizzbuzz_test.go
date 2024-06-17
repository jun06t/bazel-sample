package fizzbuzz

import (
	"os"
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	goldenFile := "testdata/golden.txt"
	expectedOutput, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("failed reading golden file: %s", err)
	}

	expectedLines := strings.Split(string(expectedOutput), "\n")
	for i := 1; i < len(expectedLines); i++ {
		expected := expectedLines[i-1]
		actual := FizzBuzz(i)
		if actual != expected {
			t.Errorf("FizzBuzz(%d) = %s; want %s", i, actual, expected)
		}
	}
}
