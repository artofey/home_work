package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type cmdCases struct {
	inCMD  []string
	outStr string
}

type envCases struct {
	inEnv       Environment
	expectedEnv Environment
}

func getCurrentEnvs(envs Environment) Environment {
	currentEnvs := make(Environment)
	for nameEnv := range envs {
		if v, ok := os.LookupEnv(nameEnv); ok {
			currentEnvs[nameEnv] = v
		}
	}
	return currentEnvs
}

func TestRunCmd(t *testing.T) {
	// t.Skip()
	t.Run("test envs", func(t *testing.T) {
		// Установка плохих значений.
		os.Setenv("HELLO", "SHOULD_REPLACE")
		os.Setenv("FOO", "SHOULD_REPLACE")
		os.Setenv("UNSET", "SHOULD_REMOVE")

		tests := []envCases{
			{
				inEnv: Environment{
					"BAR":   "bar",
					"FOO":   "  foo\nwith new line",
					"UNSET": "",
				},
				expectedEnv: Environment{
					"BAR": "bar",
					"FOO": "  foo\nwith new line",
				},
			},
			{
				inEnv:       Environment{},
				expectedEnv: Environment{},
			},
		}
		for i, test := range tests {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				cmd := []string{"pwd"}
				code := RunCmd(cmd, test.inEnv)
				require.Equal(t, 0, code)
				require.Equal(t, test.expectedEnv, getCurrentEnvs(test.inEnv))
			})
		}
	})

	t.Run("without args", func(t *testing.T) {
		code := RunCmd([]string{}, Environment{})
		require.Equal(t, 1, code)

	})
}
