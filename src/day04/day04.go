package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//day04_1()
	day04_2()
}

var target = "XMAS"
var target2 = "SAMX"

func checkEquals(s string) bool {
	return s == target || s == target2
}

func findString(i, j int, h []string) int {
	res := 0
	leftDiagonal, rightDiagonal, right, bottom := "", "", "", ""
	if len(h[i]) > j+3 {
		right = h[i][j : j+4]
	}

	if len(h) > i+3 {
		bottom = string(h[i][j]) + string(h[i+1][j]) + string(h[i+2][j]) + string(h[i+3][j])
	}
	if len(h) > i+3 && len(h[i+3]) > j+3 {
		rightDiagonal = string(h[i][j]) + string(h[i+1][j+1]) + string(h[i+2][j+2]) + string(h[i+3][j+3])
	}
	if len(h) > i+3 && j-3 >= 0 {
		leftDiagonal = string(h[i][j]) + string(h[i+1][j-1]) + string(h[i+2][j-2]) + string(h[i+3][j-3])
	}

	if checkEquals(bottom) {
		res++
	}
	if checkEquals(leftDiagonal) {
		res++
	}
	if checkEquals(right) {
		res++
	}
	if checkEquals(rightDiagonal) {
		res++
	}

	return res
}

func findX_Mas(lines []string) int {
	res := 0

	window := make([]string, 3)

	for i := 0; i < len(lines)-2; i++ {
		for j := 0; j < len(lines[i])-2; j++ {
			copy(window, lines[i:i+3])
			for k := 0; k < 3; k++ {
				window[k] = window[k][j : j+3]
			}
			masCount := findMas(window)
			if masCount == 2 {
				res++
			}
		}
	}
	return res
}

func findMas(h []string) int {
	count := 0

	t := "MAS"
	t2 := "SAM"

	right := string(h[0][0]) + string(h[1][1]) + string(h[2][2])
	left := string(h[0][2]) + string(h[1][1]) + string(h[2][0])

	if right == t || right == t2 {
		count++
	}
	if left == t || left == t2 {
		count++
	}
	return count
}

func day04_1() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}
	res := 0

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		for j := range line {
			res += findString(i, j, lines)
		}
	}

	fmt.Println(res)
}

func day04_2() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}
	res := 0

	lines := strings.Split(string(input), "\n")

	res = findX_Mas(lines)

	fmt.Println("Output is: ", res)

}
