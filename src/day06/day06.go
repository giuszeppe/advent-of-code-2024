package main

import (
	"bytes"
	"fmt"
	"os"
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
	res := 1 // we include the starting position
	visited := make(map[[2]int]struct{})
	// i create a map of [x,y] coordinates to see if visited
	// loop through all to find the caret
	// when creating the array, store the position of the caret

	m := bytes.Split(input, []byte{'\n'})
	caretPos := [2]int{}
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[row]); col++ {
			if m[row][col] == '^' {
				caretPos[0], caretPos[1] = row, col
			}
		}
	}

	xPos, yPos := caretPos[0], caretPos[1]
	for {
		// handle different positions of the caret
		caret := m[xPos][yPos]
		nextPos := [2]int{}

		switch caret {
		case '>':
			nextPos = [2]int{xPos, yPos + 1}
		case '<':
			nextPos = [2]int{xPos, yPos - 1}
		case '^':
			nextPos = [2]int{xPos - 1, yPos}
		case 'v':
			nextPos = [2]int{xPos + 1, yPos}
		}

		if nextPos[0] >= len(m) || nextPos[0] < 0 || nextPos[1] < 0 || nextPos[1] >= len(m[0]) {
			break
		}

		if m[nextPos[0]][nextPos[1]] == '#' {
			m[xPos][yPos] = rotateCaret90ToTheRight(caret)
			continue
		}

		_, ok := visited[nextPos]
		if !ok {
			res++
			visited[nextPos] = struct{}{}
		}

		m[xPos][yPos] = '.'
		xPos, yPos = nextPos[0], nextPos[1]
		m[xPos][yPos] = caret
	}

	fmt.Println(res)
}

func rotateCaret90ToTheRight(caret byte) byte {
	switch caret {

	case '>':
		return 'v'
	case '<':
		return '^'
	case '^':
		return '>'
	case 'v':
		return '<'
	}
	return caret
}

type Pos struct {
	x int
	y int
}
type Point struct {
	x   int
	y   int
	dir byte
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

	mOriginal := bytes.Split(input, []byte{'\n'})
	caretPos := Pos{}
	for row := 0; row < len(mOriginal); row++ {
		for col := 0; col < len(mOriginal[row]); col++ {
			if mOriginal[row][col] == '^' {
				caretPos = Pos{x: row, y: col}
			}
		}
	}

	for i := 0; i < len(mOriginal); i++ {
		for j := 0; j < len(mOriginal[i]); j++ {
			if mOriginal[i][j] == '#' || mOriginal[i][j] == '^' {
				continue
			}
			m := make([][]byte, len(mOriginal))
			for k := range mOriginal {
				m[k] = make([]byte, len(mOriginal[k]))
				copy(m[k], mOriginal[k])
			}

			m[i][j] = 'O'
			visited := make(map[Point]struct{})

			xPos, yPos := caretPos.x, caretPos.y

			for {
				// handle different positions of the caret
				//fmt.Println(xPos, yPos)
				caret := m[xPos][yPos]
				nextPos := Pos{}
				curPoint := Point{x: xPos, y: yPos, dir: caret}

				if _, ok := visited[curPoint]; !ok {
					visited[curPoint] = struct{}{}
				} else {
					res++
					break
				}

				switch caret {
				case '>':
					nextPos = Pos{x: xPos, y: yPos + 1}
				case '<':
					nextPos = Pos{x: xPos, y: yPos - 1}
				case '^':
					nextPos = Pos{x: xPos - 1, y: yPos}
				case 'v':
					nextPos = Pos{x: xPos + 1, y: yPos}
				}

				if nextPos.x >= len(m) || nextPos.x < 0 || nextPos.y < 0 || nextPos.y >= len(m[0]) {
					break
				}

				if m[nextPos.x][nextPos.y] == '#' || m[nextPos.x][nextPos.y] == 'O' {
					m[xPos][yPos] = rotateCaret90ToTheRight(caret)
					continue
				}
				m[xPos][yPos] = '.'
				xPos, yPos = nextPos.x, nextPos.y
				m[xPos][yPos] = caret
			}

			m[i][j] = '.'
		}
	}
	fmt.Println(res)
}
