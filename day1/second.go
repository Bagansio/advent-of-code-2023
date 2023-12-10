package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var englishNumberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func extractFirst(line string) string {

	for len(line) > 0 {
		if line[0] >= '0' && line[0] <= '9' {
			return string(line[0])
		}

		for word, value := range englishNumberMap {
			if strings.HasPrefix(line, word) {
				return value
			}
		}
		line = line[1:]
	}
	return "0"
}

func extractLast(line string) string {

	for len(line) > 0 {
		lastIndex := len(line) - 1
		if line[lastIndex] >= '0' && line[lastIndex] <= '9' {
			return string(line[lastIndex])
		}

		for word, value := range englishNumberMap {
			if strings.HasSuffix(line, word) {
				return value
			}
		}
		line = line[:lastIndex]
	}
	return "0"
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
		first := extractFirst(line)
		last := extractLast(line)

		num, _ := strconv.Atoi(first + last)
		sum += num

	}
	fmt.Println(sum)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
