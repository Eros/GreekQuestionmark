package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func replace(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ";") {
			strings.ReplaceAll(scanner.Text(), ";", ";")
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}
}

func main() {

	if os.Args[0] == "" {
		return
	}

	replace(os.Args[0])
}
