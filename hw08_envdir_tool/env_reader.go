package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// Environment is map of OS environments.
type Environment map[string]string

func prepareEnvVal(s string) string {
	// Удалить пробельные символы справа.
	text := strings.TrimRightFunc(s, unicode.IsSpace)
	// Удалить двойные кавычки.
	// text = strings.TrimFunc(text, func(r rune) bool { return r == '"' })
	// Зафменить терминальные нули на перевод строки.
	return string(bytes.ReplaceAll([]byte(text), []byte{'\x00'}, []byte("\n")))
}

func getEnvValue(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("opening file %v raised error: %w", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	return prepareEnvVal(scanner.Text()), nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envs := make(Environment)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v raised error: %w", dir, err)
	}

	for _, file := range files {
		if file.IsDir() || strings.ContainsRune(file.Name(), '=') {
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
