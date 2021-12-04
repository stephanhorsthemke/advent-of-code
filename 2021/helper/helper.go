package helper

import (
	"bufio"
	"log"
	"os"
)



func GetInput(file string) []string {
	readFile, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	defer readFile.Close()

	return fileTextLines
}