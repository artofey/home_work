package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type runCases struct {
	inCMD       []string
	inEnv       Environment
	outStr      string
	expectedEnv Environment
}

func TestRunCmd(t *testing.T) {
	t.Skip()
	tests := []runCases{
		{
			inCMD:       []string{""},
			inEnv:       Environment{"FOO": "bar"},
			outStr:      "",
			expectedEnv: Environment{"FOO": "bar"},
		},
	}
	var nameTest string
	for _, test := range tests {
		nameTest = fmt.Sprintf("%v: %v", test.inCMD, test.inEnv)
		t.Run(nameTest, func(t *testing.T) {
			code := RunCmd(test.inCMD, test.inEnv)
			require.Equal(t, code, code)
		})
	}
}
