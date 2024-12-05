package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	l1 := []int{}
	l2 := []int{}
	m := make(map[int]int)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "   ")
		n1, err := strconv.Atoi(a[0])
		n2, err2 := strconv.Atoi(a[1])
		if err != nil || err2 != nil {
			panic(err)
		}
		l1, l2 = append(l1, n1), append(l2, n2)
		if _, s := m[n2]; !s {
			m[n2] = 0
		}
		m[n2]++

	}

	sort.Ints(l1)
	sort.Ints(l2)
	res := 0

	for i, n := range l1 {
		fmt.Println(i, n, m[n])
		t, _ := m[n]

		res += n * t
	}

	fmt.Println(res)
}
