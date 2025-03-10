package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	allLevels := loadAllLevelsFromFile("day2/input.txt")

	safeCount := getCountSafeLevels(allLevels)

	fmt.Printf("Safe levels: %d\n", safeCount)
}

func getCountSafeLevels(allLevels [][]int) int {
	var safeCount int
	for _, levels := range allLevels {
		if areLevelsSafeWithRemoval(levels) {
			safeCount++
		}
	}

	return safeCount
}

func areLevelsSafeWithRemoval(levels []int) bool {
	fmt.Printf("Checking levels: %v ", levels)

	isAscending := levels[0] < levels[1]
	hasIgnoredLevel := false

	for i := 0; i < len(levels)-1; i++ {

		if !isOrderSafe(levels[i], levels[i+1], isAscending) || !isDiffSafe(levels[i], levels[i+1]) {

			if hasIgnoredLevel {
				fmt.Println("Unsafe: have already removed a level")
				return false
			}

			indexToRemove := testRemoval(levels)
			if indexToRemove == -1 {
				fmt.Println("Unsafe: No removal will fix error")
				return false
			}

			upper := indexToRemove + 1
			if upper > len(levels)-1 {
				levels = levels[:len(levels)-1]
			} else {
				levels = append(levels[:indexToRemove], levels[upper:]...)
			}
			isAscending = levels[0] < levels[1]
			fmt.Printf("new levels %v\n", levels)
			i = 0
			hasIgnoredLevel = true
		}
	}

	fmt.Printf("Safe\n")
	return true
}

func areLevelsSafe(levels []int) bool {
	fmt.Printf("Checking: %v ", levels)

	ascending := levels[0] < levels[1]

	for i := 0; i < len(levels)-1; i++ {

		if !isOrderSafe(levels[i], levels[i+1], ascending) || !isDiffSafe(levels[i], levels[i+1]) {
			fmt.Println("Unsafe")
			return false
		}
	}

	fmt.Printf("Safe\n")
	return true
}

func testRemoval(levels []int) int {

	for i := 0; i < len(levels); i++ {
		cloneLevels := make([]int, len(levels))
		copy(cloneLevels, levels)
		if i == len(levels)-1 {
			if areLevelsSafe(cloneLevels[:len(cloneLevels)-1]) {
				return i
			}
		} else {
			if areLevelsSafe(append(cloneLevels[:i], cloneLevels[i+1:]...)) {
				return i
			}
		}
	}

	return -1
}

func isOrderSafe(i1 int, i2 int, isAscending bool) bool {
	return (i1 < i2) == isAscending
}

func isDiffSafe(i1 int, i2 int) bool {
	diff := math.Abs(float64(i1 - i2))
	return diff <= 3 && diff >= 1
}

func loadAllLevelsFromFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var allLevels [][]int

	for scanner.Scan() {
		line := scanner.Text()

		stringSlice := strings.Split(line, " ")
		levels := []int{}

		for _, s := range stringSlice {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			levels = append(levels, i)
		}

		fmt.Printf("Levels loaded: %v\n", levels)
		allLevels = append(allLevels, levels)
	}

	return allLevels
}
