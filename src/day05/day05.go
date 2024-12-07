package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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
	res := 0

	lines := strings.Split(string(input), "A")
	rules := strings.Split(lines[0], "\n")
	updates := strings.Split(lines[1], "\n")

	mustOccurBefore := map[string][]string{}

	for _, rule := range rules {
		pages := strings.Split(rule, "|")
		if len(pages) == 1 {
			continue
		}
		x := pages[0] // must occur before y
		y := pages[1] // must occur after x
		mustOccurBefore[y] = append(mustOccurBefore[y], x)
	}

	for _, update := range updates {
		f := true
		pages := strings.Split(update, ",")
		if len(pages) == 0 {
			continue
		}

		for i := 0; i < len(pages); i++ {
			mustBefore := mustOccurBefore[pages[i]]

			// check if any page after the one i'm looking at must happen before the one i'm checking
			for j := i; j < len(pages); j++ {
				if slices.Contains(mustBefore, pages[j]) {
					f = false
					break
				}
			}
			if !f {
				break
			}

		}

		if f && len(pages) != 0 {
			n, _ := strconv.Atoi(string(pages[len(pages)/2]))
			res += n
		}

	}
	fmt.Println(res)

}

func day2() {
	input, err := os.ReadFile("./input1.txt")
	if err != nil {
		panic(err)
	}
	res := 0

	lines := strings.Split(string(input), "A")
	rules := strings.Split(lines[0], "\n")
	updates := strings.Split(lines[1], "\n")
	wrongUpdates := []string{}

	mustOccurBefore := map[string][]string{}
	mustOccurAfter := map[string][]string{}

	for _, rule := range rules {
		pages := strings.Split(rule, "|")
		if len(pages) == 1 {
			continue
		}
		x := pages[0] // must occur before y
		y := pages[1] // must occur after x
		mustOccurBefore[y] = append(mustOccurBefore[y], x)
		mustOccurAfter[x] = append(mustOccurAfter[x], y)
	}

	for _, update := range updates {
		f := true
		pages := strings.Split(update, ",")
		if len(pages) == 0 {
			continue
		}

		for i := 0; i < len(pages); i++ {
			mustBefore := mustOccurBefore[pages[i]]

			// check if any page after the one i'm looking at must happen before the one i'm checking
			for j := i; j < len(pages); j++ {
				if slices.Contains(mustBefore, pages[j]) {
					f = false
					break
				}
			}
			if !f {
				break
			}
		}
		if !f && len(pages) != 0 {
			wrongUpdates = append(wrongUpdates, update)
			fixedUpdate := fixUpdate(update, mustOccurBefore)
			fmt.Println(update, fixedUpdate)
			n, _ := strconv.Atoi(fixedUpdate[len(fixedUpdate)/2])
			res += n
		}
	}
	fmt.Println(res)

}

func fixUpdate(update string, mustOccurAfter map[string][]string) (fixedUpdate []string) {
	// 75,97,47,61,53
	// 97,75,47,61,53
	candidates := strings.Split(update, ",")
	canPlace := true

	// cycle over every element in the se
	// for every element remaining, check if the element i want to place is in the mustOccurAfter set
	// if so, return false. This is not the correct index to place the element.
	// if it is, return true and remove the element
	// at the end of each iteration, one element is placed in the correct place

	// ensures eventually every elem gets placed
	for len(candidates) > 0 {
		fmt.Println((candidates))
		// cycle through all the elements
		for i, elem := range candidates {
			canPlace = true
			for _, comparison := range candidates {
				if elem == comparison {
					continue
				}

				if slices.Contains(mustOccurAfter[comparison], elem) {
					canPlace = false
					break
				}
			}

			if canPlace {
				fixedUpdate = append(fixedUpdate, elem)
				candidates = append(candidates[:i], candidates[i+1:]...)
				break
			}
		}

	}
	return fixedUpdate
}
