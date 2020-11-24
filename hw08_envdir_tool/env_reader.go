package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func prepareEnvVal(b []byte) EnvValue {
	var ev EnvValue
	// Удалить пробельные символы справа.
	b = bytes.TrimRightFunc(b, unicode.IsSpace)
	// Зафменить терминальные нули на перевод строки.
	b = bytes.ReplaceAll(b, []byte{'\x00'}, []byte("\n"))
	if len(b) == 0 {
		ev.NeedRemove = true
	}
	ev.Value = string(b)
	return ev
}

func getEnvValue(fileName string) (EnvValue, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return EnvValue{}, fmt.Errorf("opening file %v raised error: %w", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return EnvValue{}, fmt.Errorf("reading file %v raised error: %w", fileName, err)
	}
	return prepareEnvVal(scanner.Bytes()), nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v raised error: %w", dir, err)
	}
	envs := make(Environment, len(files))

	for _, file := range files {
		if file.IsDir() || strings.ContainsRune(file.Name(), '=') {
			continue
		}
		envVal, err := getEnvValue(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		envs[file.Name()] = envVal
	}
	return envs, nil
}
