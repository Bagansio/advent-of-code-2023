package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// mapMapping represents a single map (e.g., seed-to-soil map)
type mapMapping []mappingEntry

// mappingEntry represents a single entry in a map
type mappingEntry struct {
	destStart   int
	sourceStart int
	length      int
}

// readInputFromFile reads seeds and maps from a file and returns them
func readInputFromFile(filename string) ([]int, []mapMapping) {
	var seeds []int
	var maps []mapMapping

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds:") {
			// Read seeds from the line starting with "seeds:"
			seeds = readSeeds(line)
		} else {
			maps = append(maps, readMap(scanner))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return seeds, maps
}

// readSeeds reads seeds from a line and returns them as a slice
func readSeeds(line string) []int {
	fields := strings.Fields(line)
	var seeds []int
	for _, field := range fields[1:] {
		seed, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Error converting seed:", err)
			os.Exit(1)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

// readMapsFromFile reads the maps from a file and returns a slice of maps
func readMapsFromFile(filename string) []mapMapping {
	var maps []mapMapping

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seed-to-soil map:") {
			// Read the map for seed-to-soil and append to the list
			maps = append(maps, readMap(scanner))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return maps
}

// readMap reads a single map from the input and returns it as a mapMapping
func readMap(scanner *bufio.Scanner) mapMapping {
	var mapping mapMapping

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var destStart, sourceStart, length int
		fmt.Sscanf(line, "%d %d %d", &destStart, &sourceStart, &length)
		mapping = append(mapping, mappingEntry{destStart, sourceStart, length})
	}

	return mapping
}

// applyMappings applies the mappings to the given seed and returns the resulting value
func applyMappings(seed int, currentMap map[int]int, maps []mapMapping) int {
	for _, m := range maps {
		for _, entry := range m {
			if seed >= entry.sourceStart && seed < entry.sourceStart+entry.length {
				// Update the seed and continue applying mappings
				seed = entry.destStart + (seed - entry.sourceStart)
				break
			}
		}
	}
	// If no mapping is found, return the seed itself
	return seed
}

// findLowestLocation finds the lowest location number from the current map
func findLowestLocation(currentMap map[int]int) int {
	lowest := -1
	for _, location := range currentMap {
		if lowest == -1 || location < lowest {
			lowest = location
		}
	}
	return lowest
}

func main() {
	// Read input data from a file
	seeds, maps := readInputFromFile("input.txt")

	// Initialize a map to store the current mapping
	currentMap := make(map[int]int)

	// Iterate through each seed and apply the mappings
	for _, seed := range seeds {
		currentMap[seed] = applyMappings(seed, currentMap, maps)
	}

	// Find the lowest location number
	lowestLocation := findLowestLocation(currentMap)

	// Print the result
	fmt.Println("Lowest Location Number:", lowestLocation)
}
