package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
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
		err := os.Unsetenv(key)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Setenv(key, val)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = c.Wait()
	// fmt.Println(os.Environ())
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// fmt.Println(exitErr.Error())
			code, _ := strconv.Atoi(exitErr.Error())
			return code
		}
	}
	return 0
}
