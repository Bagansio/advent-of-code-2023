package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const RED_MAX = 12
const BLUE_MAX = 14
const GREEN_MAX = 13

var cubes = map[string]int{
	"red":   RED_MAX,
	"blue":  BLUE_MAX,
	"green": GREEN_MAX,
}

func checkIfPossible(line string) int {
	list := strings.Split(line, " ")
	gameFragment := list[1]
	gameId, _ := strconv.Atoi(gameFragment[:len(gameFragment)-1])

	list = strings.Split(line, ":")
	gameSets := strings.Split(list[1], ";")

	for _, gameSet := range gameSets {
		combinations := strings.Split(gameSet, ",")

		for _, combination := range combinations {
			values := strings.Split(combination, " ")

			value, _ := strconv.Atoi(values[1])
			if value > cubes[values[2]] {
				return 0
			}
		}
	}
	return gameId
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
	sumIds := 0
	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Call the extractNumbers function for each line
		sumIds += checkIfPossible(line)

	}
	fmt.Println(sumIds)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
