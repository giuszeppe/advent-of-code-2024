package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	day03_1()
	day03_2()
}

func day03_1() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	res := 0

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
		ss := r.FindAllStringSubmatch(scanner.Text(), -1)
		for _, group := range ss {
			n1, _ := strconv.Atoi(group[1])
			n2, _ := strconv.Atoi(group[2])
			res += n1 * n2
		}
	}
	fmt.Println(res)
}

func day03_2() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	res := 0

	scanner := bufio.NewScanner(f)
	calculate := true
	for scanner.Scan() {
		r, _ := regexp.Compile(`(don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\))`)
		ss := r.FindAllStringSubmatch(scanner.Text(), -1)
		//fmt.Println(ss)

		for _, group := range ss {

			if v, _ := regexp.Match(`do\(\)`, []byte(group[0])); v {
				calculate = true
				continue
			}
			if v, _ := regexp.Match(`don't\(\)`, []byte(group[0])); v {
				calculate = false
				continue
			}
			if calculate {
				n1, _ := strconv.Atoi(group[2])
				n2, _ := strconv.Atoi(group[3])
				res += n1 * n2
			}
		}
	}
	fmt.Println(res)
}
