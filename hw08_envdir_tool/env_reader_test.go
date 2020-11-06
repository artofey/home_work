package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	input    string
	expected Environment
	err      interface{}
}

func TestReadDir(t *testing.T) {
	testCases := []testCase{
		{
			input:    "testdata/nodir",
			expected: nil,
			err:      &os.PathError{},
		},
		{
			input: "testdata/env",
			expected: Environment{
				"BAR":   "bar",
				"HELLO": "hello",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			res, err := ReadDir(test.input)

			require.Equal(t, test.expected, res)
			require.IsType(t, test.err, err)
		})
	}
}
