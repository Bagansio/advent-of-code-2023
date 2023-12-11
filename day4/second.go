package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkIfPossible(line string, extras *[]int, index int) int {
	list := strings.Split(line, ":")
	numbers := strings.Split(list[1], "|")

	winningList := strings.Fields(numbers[0])
	winningMap := make(map[string]bool)
	for _, elem := range winningList {
		winningMap[elem] = false
	}

	ownList := strings.Fields(numbers[1])

	for _, elem := range ownList {
		_, ok := winningMap[elem]

		if ok {
			winningMap[elem] = true
		}
	}

	sum := 0

	for _, match := range winningMap {
		if match {
			sum++
		}
	}

	value := (*extras)[index] + 1

	for i := index + 1; i <= index+sum; i++ {
		if i >= len(*extras) {
			*extras = append(*extras, value)
		} else {
			(*extras)[i] += value
		}
	}

	return value

}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)
	sum := 0
	index := 0
	extras := []int{}
	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		if len(extras) <= index {
			extras = append(extras, 0)
		}

		// Call the extractNumbers function for each line
		sum += checkIfPossible(line, &extras, index)
		index++
	}
	fmt.Println(sum)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
