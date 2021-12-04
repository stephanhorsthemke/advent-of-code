package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	input := getInput("input.txt")
	numbers := make([]int, len(input))
	for i, v := range input {
		numbers[i], _ = strconv.Atoi(v)
	}

	for _, f := range numbers {
		for _, s := range numbers {
			for _, t := range numbers {
				if f+s+t == 2020 {
					fmt.Println(f * s * t)
					return
				}
			}
		}
	}

}

func getInput(file string) []string {
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
