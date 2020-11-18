package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	input    string
	expected Environment
	err      error
}

func errCheck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func errCheckExist(e error) {
	if e != nil && !errors.Is(e, os.ErrExist) {
		panic(e)
	}
}

// Создать данные необходимые для тестов
func makeTestData() {
	td := "testdata/envequal"
	err := os.Mkdir(td, 0777)
	errCheckExist(err)

	err = os.Chdir(td)
	errCheck(err)
	cases := map[string]string{
		"=HELLO": "1",
		"H=ELLO": "2",
		"HELLO":  "3",
		"HELLO=": "4",
	}
	for fileName, val := range cases {
		err = ioutil.WriteFile(fileName, []byte(val), 0777)
		errCheckExist(err)
	}
	err = os.Chdir("../..")
	errCheck(err)
}

func removeTestData() {
	err := os.RemoveAll("testdata/envequal")
	errCheck(err)
}

func TestReadDir(t *testing.T) {
	t.Run("if dir not exist", func(t *testing.T) {
		_, err := ReadDir("testdata/nodir")
		var pe *os.PathError
		require.True(t, errors.As(err, &pe))
	})

	makeTestData()
	testCases := []testCase{
		{
			input: "testdata/envequal",
			expected: Environment{
				"HELLO": "3",
			},
		},
		{
			input: "testdata/env",
			expected: Environment{
				"BAR":   "bar",
				"HELLO": "\"hello\"",
				"UNSET": "",
				"FOO": `   foo
with new line`,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			res, err := ReadDir(test.input)

			require.Equal(t, test.expected, res)
			require.True(t, errors.Is(test.err, err))
			require.NoError(t, err)
		})
	}
	removeTestData()
}

func TestPrepareEnvVal(t *testing.T) {
	testCases := map[string]string{
		"right_tab\t":  "right_tab",
		"right_space ": "right_space",
		"simple":       "simple",
		`"tests"`:      `"tests"`,
		"":             "",
		"111\x00111":   "111\n111",
	}

	for input, expected := range testCases {

		t.Run(input, func(t *testing.T) {
			result := prepareEnvVal(input)
			require.Equal(t, expected, result)
		})
	}
}
