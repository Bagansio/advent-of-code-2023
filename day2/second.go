package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMinimumSetOfCubes(line string) int {
	list := strings.Split(line, " ")
	list = strings.Split(line, ":")
	gameSets := strings.Split(list[1], ";")

	var cubes = map[string]int{
		"red":   1,
		"blue":  1,
		"green": 1,
	}

	for _, gameSet := range gameSets {
		combinations := strings.Split(gameSet, ",")

		for _, combination := range combinations {
			values := strings.Split(combination, " ")

			value, _ := strconv.Atoi(values[1])
			fmt.Println(values, values[1], value)
			if value > cubes[values[2]] {
				cubes[values[2]] = value
			}
		}
	}
	result := 1

	for _, value := range cubes {
		result *= value
	}

	return result
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
	sumSets := 0
	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Call the extractNumbers function for each line
		sumSets += getMinimumSetOfCubes(line)

	}
	fmt.Println(sumSets)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
