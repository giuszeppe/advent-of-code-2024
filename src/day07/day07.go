package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/giuszeppe/advent-of-code-2024/dsa"
)

func main() {
	day1()
	day2()
}

func day1() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}

	res := 0 // we include the starting position
	lines := strings.Split(string(input), "\n")
	allowed_operators := []string{"+", "*"}
	for _, line := range lines {
		test, operators := strings.Split(line, ":")[0], strings.Split(line, ":")[1]
		testValue, _ := strconv.Atoi(test)

		numbers := []int{}
		for _, n := range strings.Split(strings.Trim(operators, " "), " ") {
			intN, _ := strconv.Atoi(n)
			numbers = append(numbers, intN)
		}

		length := len(numbers) - 1
		calcStack := dsa.Deque[int]{}
		for _, ops := range GenerateCombinations(allowed_operators, length) {
			opIdx := 0
			for _, n := range numbers {
				if calcStack.Len() == 1 {
					n2, _ := calcStack.PopBack()
					n = performOperation(ops[opIdx], n, n2)
					opIdx++
				}
				calcStack.PushBack(n)
			}
			result, _ := calcStack.PopBack()
			if result == testValue {
				res += testValue
				break
			}
		}
	}

	fmt.Println(res)
}

func performOperation(op string, n1, n2 int) int {
	switch op {
	case "+":
		return n1 + n2
	case "*":
		return n1 * n2
	case "||":
		s1 := strconv.Itoa(n1)
		s2 := strconv.Itoa(n2)
		n3, _ := strconv.Atoi(s1 + s2)
		return n3
	}
	return 0
}

func GenerateCombinations(symbols []string, length int) [][]string {
	// Base case: if length is 0, return an empty combination
	if length == 0 {
		return [][]string{{}}
	}

	// Recursive case: build combinations by appending each symbol
	combinations := [][]string{}
	smallerCombinations := GenerateCombinations(symbols, length-1)
	for _, combo := range smallerCombinations {
		for _, symbol := range symbols {
			newCombo := append([]string{}, combo...) // Copy the existing combo
			newCombo = append(newCombo, symbol)      // Append the current symbol
			combinations = append(combinations, newCombo)
		}
	}

	return combinations
}

// I fix one position to be an obstacle
// I save the last 3 position before I hit the new obstacle
// If the first obstacle I hit after hitting the new obstacle is the firs of that array of three I have a loop
func day2() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}

	res := 0 // we include the starting position
	lines := strings.Split(string(input), "\n")
	allowed_operators := []string{"+", "*", "||"}
	for _, line := range lines {
		test, operators := strings.Split(line, ":")[0], strings.Split(line, ":")[1]
		testValue, _ := strconv.Atoi(test)

		numbers := []int{}
		for _, n := range strings.Split(strings.Trim(operators, " "), " ") {
			intN, _ := strconv.Atoi(n)
			numbers = append(numbers, intN)
		}

		length := len(numbers) - 1
		calcStack := dsa.Deque[int]{}
		for _, ops := range GenerateCombinations(allowed_operators, length) {
			opIdx := 0
			// it is n2 because number is the number after the one put into the stack
			for _, n2 := range numbers {
				if calcStack.Len() == 1 {
					n, _ := calcStack.PopBack()
					n2 = performOperation(ops[opIdx], n, n2)
					opIdx++
				}
				calcStack.PushBack(n2)
			}
			result, _ := calcStack.PopBack()
			if result == testValue {
				res += testValue
				break
			}
		}
	}

	fmt.Println(res)
}
