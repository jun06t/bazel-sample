// +build integration

package dice

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	var err error
	cli, err = newClient(context.TODO(), "localhost:6379")
	if err != nil {
		panic(err)
	}
	m.Run()
	err = cli.Close()
	if err != nil {
		panic(err)
	}
}

func TestSaveNumber(t *testing.T) {
	tests := []struct {
		name     string
		in       int
		expected string
	}{
		{
			name:     "success",
			in:       5,
			expected: "5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field, err := SaveNumber(context.TODO(), tt.in)
			require.NoError(t, err)
			out, err := cli.HGet(context.TODO(), historyKey, field).Result()
			require.NoError(t, err)
			assert.Equal(t, tt.expected, out)
		})
	}
}
