package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

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
		return "", err
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
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() || strings.ContainsRune(file.Name(), '=') {
			continue
		}
		envVal, err := getEnvValue(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		if envVal != "" {
			envs[file.Name()] = envVal
		}
	}
	return envs, nil
}
