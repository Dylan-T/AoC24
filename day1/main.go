package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 1")

	// Load input
	var l1, l2 = loadInputToSortedLists("day1/input.txt")

	// Sort input
	sort.Ints(l1)
	sort.Ints(l2)

	// loop through and sum distances
	sum := sumListDistances(l1, l2)
	fmt.Println("Sum is", sum)

	similarity := getListSimilarity(l1, l2)
	fmt.Println("Similarity is", similarity)
}

func sumListDistances(l1 []int, l2 []int) int {
	sum := 0
	for i := 0; i < len(l1); i++ {
		diff := l1[i] - l2[i]

		sum += int(math.Abs(float64(diff)))
	}

	return sum
}

func getListSimilarity(l1 []int, l2 []int) int {
	similarity := 0

	rCounter := make(map[int]int)

	for i := 0; i < len(l2); i++ {
		rCounter[l2[i]] = rCounter[l2[i]] + 1
	}

	for i := 0; i < len(l1); i++ {
		similarity += l1[i] * rCounter[l1[i]]
	}

	return similarity
}

func loadInputToSortedLists(path string) ([]int, []int) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "   ")

		val1, _ := strconv.Atoi(vals[0])
		val2, _ := strconv.Atoi(vals[1])

		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	return list1, list2
}
