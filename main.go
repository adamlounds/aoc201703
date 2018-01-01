package main

import (
	"bufio"
	//"errors"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

type position struct {
	x int
	y int
}

type vector struct {
	x int
	y int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("enter a number: ")
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("bad line: %s\n", err)
			fmt.Print("enter a number: ")
			continue
		}

		pos := findPosition(int(num))
		offset := calcOffset(pos)
		fmt.Printf("offset for %s is %d\n", line, offset)

		nextNum := spiralEntryExceeding(int(num))
		fmt.Printf("exceeding number over %s is %d\n", line, nextNum)
		fmt.Print("enter a number: ")
	}

	if scanner.Err() != nil {
		os.Stderr.WriteString(fmt.Sprintf("scan error %s", scanner.Err))
	}
}

func calcOffset(pos position) int {
	return abs(pos.x) + abs(pos.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findPosition(num int) position {
	pos := position{0, 0}
	if num == 1 {
		return pos
	}

	edgeLen := 1
	directions := map[string]vector{
		"L": vector{-1, 0},
		"R": vector{1, 0},
		"U": vector{0, -1},
		"D": vector{0, 1},
	}
	nextDirection := map[string]string{
		"R": "U",
		"U": "L",
		"L": "D",
		"D": "R",
	}
	i := int(1)
	currentDirection := "R"
findPos:
	for i <= num {
		for moveStep := 1; moveStep <= edgeLen; moveStep++ {
			i++
			vector := directions[currentDirection]
			pos.x += vector.x
			pos.y += vector.y
			//os.Stderr.WriteString(fmt.Sprintf("%d: step %d/%d %s (%v)\n", i, moveStep, edgeLen, currentDirection, pos))
			if i == num {
				//os.Stderr.WriteString(fmt.Sprintf("%d: %v\n", num, pos))
				break findPos
			}
		}
		currentDirection = nextDirection[currentDirection]
		if currentDirection == "L" {
			edgeLen++ // we go left one more step than we went up
		}
		if currentDirection == "R" {
			edgeLen++ // we go right one more step than we went down
		}
		//os.Stderr.WriteString(fmt.Sprintf("%d: end of side, will move %s next\n", i, currentDirection))
	}
	//os.Stderr.WriteString(fmt.Sprintf("------------ %v\n", pos))

	return pos
}

func spiralEntryExceeding(valueToExceed int) int {
	size := 101
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	edgeLen := 1
	directions := map[string]vector{
		"L": vector{-1, 0},
		"R": vector{1, 0},
		"U": vector{0, -1},
		"D": vector{0, 1},
	}
	nextDirection := map[string]string{
		"R": "U",
		"U": "L",
		"L": "D",
		"D": "R",
	}

	// start in the middle. int + zero-indexes work together.
	pos := position{int(size / 2), int(size / 2)}
	matrix[pos.x][pos.y] = 1

	currentDirection := "R"
	valueOfThisCell := 1
findValue:
	for true {
		for moveStep := 1; moveStep <= edgeLen; moveStep++ {
			vector := directions[currentDirection]
			pos.x += vector.x
			pos.y += vector.y

			valueOfThisCell = sumNeighbours(matrix, pos)
			matrix[pos.y][pos.x] = valueOfThisCell

			if valueOfThisCell > valueToExceed {
				break findValue
			}
		}
		currentDirection = nextDirection[currentDirection]
		if currentDirection == "L" {
			edgeLen++ // we go left one more step than we went up
		}
		if currentDirection == "R" {
			edgeLen++ // we go right one more step than we went down
		}
		//os.Stderr.WriteString(fmt.Sprintf("%d: end of side, will move %s next\n", i, currentDirection))
	}

	return valueOfThisCell
}

func sumNeighbours(matrix [][]int, pos position) int {
	x := pos.x
	y := pos.y
	above := matrix[y-1][x-1] + matrix[y-1][x] + matrix[y-1][x+1]
	sides := matrix[y][x-1] + matrix[y][x+1]
	below := matrix[y+1][x-1] + matrix[y+1][x] + matrix[y+1][x+1]
	return above + sides + below
}
