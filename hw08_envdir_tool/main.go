package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Example for run command:\ngo-envdir /path/to/env/dir command arg1 arg2")
		return
	}
	envDir := args[0]
	command := args[1:]
	env, err := ReadDir(envDir)
	if err != nil {
		log.Fatal(err)
	}
	code := RunCmd(command, env)
	os.Exit(code)
}
