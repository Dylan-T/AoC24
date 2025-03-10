package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadMemoryToString(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	memory := ""

	for scanner.Scan() {
		memory += scanner.Text()
	}

	return memory
}

func main() {
	memory := loadMemoryToString("day3/input.txt")

	result := multiplyNumbers(memory)

	fmt.Printf("Result: %v\n", result)
}

func multiplyNumbers(memory string) int {
	var result int

	doMult := true

	cmdArray := strings.Split(memory, "mul(")
	for _, cmd := range cmdArray {

		beforeS, afterS, found := strings.Cut(cmd, ")")
		if !found {
			afterS = beforeS
		}

		if strings.Contains(afterS, "don't()") {
			doMult = false
			continue
		}
		if strings.Contains(afterS, "do()") {
			doMult = true
			continue
		}

		if !found || !doMult {
			continue
		}

		arg1, arg2, err := parseArguments(beforeS)
		if err != nil {
			continue
		}

		result += arg1 * arg2
	}
	return result
}

func parseArguments(beforeS string) (int, int, error) {
	args := strings.Split(beforeS, ",")
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("incorrect number of arguments: %v", len(args))
	}
	arg1, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, fmt.Errorf("argument 1 is not a number: %v", arg1)
	}

	arg2, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, fmt.Errorf("argument 2 is not a number: %v", arg2)
	}

	return arg1, arg2, nil
}
