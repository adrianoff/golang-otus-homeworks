package main

import (
	"fmt"
	"os"
)

func main() {
	var path, command string

	path = os.Args[1]
	command = os.Args[2]
	others := os.Args[3:]

	_, _ = ReadDir(path)
	fmt.Println(path, command, others)

	//cmd := exec.Command("git", "commit", "-am", "fix")
}
