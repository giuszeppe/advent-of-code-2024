package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	day02_1()
	day02_2()
}

func day02_1() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	res := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		level := strings.Split(scanner.Text(), " ")
		if levelIsSafe(level) {
			res++
		}

	}
	fmt.Println(res)
}
func day02_2() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	res := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		level := strings.Split(scanner.Text(), " ")
		if levelIsSafe(level) {
			res++
		} else {
			for i := 0; i < len(level); i++ {
				modLevel := make([]string, len(level))
				copy(modLevel, level)
				modLevel = remove(modLevel, i)
				if levelIsSafe(modLevel) {
					res++
					break
				}
			}
		}

	}
	fmt.Println(res)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func levelIsSafe(levels []string) bool {
	isIncreasing := false
	l1, _ := strconv.Atoi(levels[0])
	l2, _ := strconv.Atoi(levels[1])

	if l1 < l2 {
		isIncreasing = true
	}
	for i := 1; i < len(levels); i++ {
		l1, _ := strconv.Atoi(levels[i-1])
		l2, _ := strconv.Atoi(levels[i])
		diff := l2 - l1

		if diff == 0 {
			return false
		}

		if diff < -3 || diff > 3 {
			return false
		}

		if (isIncreasing && diff < 1) || (!isIncreasing && diff > -1) {
			return false
		}

	}
	return true
}
