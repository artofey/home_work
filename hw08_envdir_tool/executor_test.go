package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type envCases struct {
	inEnvs      Environment
	expectedEnv map[string]string
}

func getCurrentEnvs(envs Environment) map[string]string {
	currentEnvs := make(map[string]string)
	for nameEnv := range envs {
		if v, ok := os.LookupEnv(nameEnv); ok {
			currentEnvs[nameEnv] = v
		}
	}
	return currentEnvs
}

func TestRunCmd(t *testing.T) {
	t.Run("test envs", func(t *testing.T) {
		// Установка плохих значений.
		os.Setenv("HELLO", "SHOULD_REPLACE")
		os.Setenv("FOO", "SHOULD_REPLACE")
		os.Setenv("UNSET", "SHOULD_REMOVE")

		tests := []envCases{
			{
				inEnvs: Environment{
					"BAR":   EnvValue{"bar", false},
					"FOO":   EnvValue{"  foo\nwith new line", false},
					"UNSET": EnvValue{"", true},
				},
				expectedEnv: map[string]string{
					"BAR": "bar",
					"FOO": "  foo\nwith new line",
				},
			},
			{
				inEnvs:      Environment{},
				expectedEnv: make(map[string]string),
			},
		}
		for i, test := range tests {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				cmd := []string{"pwd"}
				code := RunCmd(cmd, test.inEnvs)
				require.Equal(t, 0, code)
				require.Equal(t, test.expectedEnv, getCurrentEnvs(test.inEnvs))
			})
		}
	})
}
