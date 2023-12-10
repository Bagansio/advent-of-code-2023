package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func readFileIntoMatrix(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

func extractParts(matrix [][]rune) int {
	sum := 0

	for i, row := range matrix {
		var currentNumber int
		var currentPositions [][2]int

		for j, char := range row {
			if unicode.IsDigit(char) {
				digitValue := int(char - '0')
				currentNumber = currentNumber*10 + digitValue
				currentPositions = append(currentPositions, [2]int{i, j})
			} else if len(currentPositions) > 0 {
				if checkPart(matrix, currentPositions) {
					sum += currentNumber
				}
				currentNumber = 0
				currentPositions = nil
			}
		}

		// Handle the case where the row ends with a number
		if len(currentPositions) > 0 {
			if checkPart(matrix, currentPositions) {
				sum += currentNumber
			}
		}
	}

	return sum
}

func checkAdjacentPositions(matrix [][]rune, position [2]int) bool {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			posX := position[0] + x
			posY := position[1] + y
			//fmt.Println(posX, posY, len(matrix[0]), len(matrix))
			if posX > 0 && posY > 0 && posY < len(matrix[0]) && posX < len(matrix) && !unicode.IsDigit(matrix[posX][posY]) && matrix[posX][posY] != '.' {
				return true
			}
		}
	}
	return false
}

func checkPart(matrix [][]rune, positions [][2]int) bool {
	for _, position := range positions {
		if checkAdjacentPositions(matrix, position) {
			return true
		}
	}

	return false
}

func main() {
	filename := "input.txt"

	matrix, err := readFileIntoMatrix(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	numbers := extractParts(matrix)

	fmt.Println(numbers)
}
