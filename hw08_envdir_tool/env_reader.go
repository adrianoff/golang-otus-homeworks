package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environment := make(Environment)

	for _, e := range entries {
		if strings.Contains(e.Name(), "=") {
			continue
		}
		line := readLine(dir + "/" + e.Name())
		needToRemove := false
		if len(line) == 0 {
			needToRemove = true
		} else {
			line = strings.ReplaceAll(line, string([]byte{0x00}), "\n")
			line = strings.TrimRight(line, " \t")
		}

		environment[e.Name()] = EnvValue{
			Value:      line,
			NeedRemove: needToRemove,
		}
	}

	return environment, nil
}

func readLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		log.Fatal(err)
	}
	if fileInfo.Size() == 0 {
		return ""
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()
	file.Close()

	return line
}
