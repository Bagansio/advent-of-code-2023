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

	symbols := make(map[[2]int][]int)
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

				pos, found := checkPart(matrix, currentPositions)
				if found {
					value, ok := symbols[pos]
					if ok {
						value = append(value, currentNumber)
						symbols[pos] = value
					} else {
						temp := []int{currentNumber}
						symbols[pos] = temp
					}
				}
				currentNumber = 0
				currentPositions = nil
			}
		}

		// Handle the case where the row ends with a number
		if len(currentPositions) > 0 {
			pos, found := checkPart(matrix, currentPositions)
			if found {
				value, ok := symbols[pos]
				if ok {
					value = append(value, currentNumber)
					symbols[pos] = value
				} else {
					temp := []int{currentNumber}
					symbols[pos] = temp
				}
			}
		}
	}
	//check gears with only 2 symbols
	for _, numbers := range symbols {
		if len(numbers) == 2 {
			sum += numbers[0] * numbers[1]
		}
	}
	return sum
}

func checkAdjacentPositions(matrix [][]rune, position [2]int) ([2]int, bool) {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			posX := position[0] + x
			posY := position[1] + y
			//fmt.Println(posX, posY, len(matrix[0]), len(matrix))
			if posX > 0 && posY > 0 && posY < len(matrix[0]) && posX < len(matrix) && matrix[posX][posY] == '*' {
				return [2]int{posX, posY}, true
			}
		}
	}
	return [2]int{}, false
}

func checkPart(matrix [][]rune, positions [][2]int) ([2]int, bool) {
	for _, position := range positions {
		pos, found := checkAdjacentPositions(matrix, position)
		if found {
			return pos, found
		}
	}

	return [2]int{}, false
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
