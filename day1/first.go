package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func extractNumbers(input string) []string {
	var numbers []string

	for _, char := range input {
		if unicode.IsDigit(char) {
			// If the character is a digit, add it to the numbers slice
			numbers = append(numbers, string(char))
		}
	}

	return numbers
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
	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Call the extractNumbers function for each line
		result := extractNumbers(line)
		first := result[0]
		last := result[len(result)-1]

		num, _ := strconv.Atoi(first + last)
		sum += num

	}
	fmt.Println(sum)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
