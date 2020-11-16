package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here
	c := exec.Command(cmd[0], cmd[0:]...)
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
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Println(exitErr.Error())
			code, _ := strconv.Atoi(exitErr.Error())
			return code
		}
	}
	return 0
}
