package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := cmd[0]
	var args []string
	args = append(args, cmd[1:]...)
	c := exec.Command(command, args...)
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	for key, val := range env {
		if val.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}
		err := os.Setenv(key, val.Value)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Wait()
	if err != nil {
		var exitErr *exec.ExitError
		if ok := errors.As(err, &exitErr); ok {
			return exitErr.ExitCode()
		}
	}
	return 0
}
