package main

import (
	"fmt"

	"../helper"
)

func main() {

	input := helper.GetInput("input.txt")
	increasedCount := 0
	for i, v := range input[1:] {
		if input[i-1] < v {
			increasedCount++
		}
	}
	fmt.Println(increasedCount)
}
