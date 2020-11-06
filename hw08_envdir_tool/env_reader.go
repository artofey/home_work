package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Environment is map of OS environments.
type Environment map[string]string

func getEnvValue(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Println(fileName, scanner.Text())
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return "", nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envs := make(Environment)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		envVal, err := getEnvValue(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		envs[file.Name()] = envVal
	}
	return envs, nil
}
