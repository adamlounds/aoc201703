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
	x int64
	y int64
}

type vector struct {
	x int64
	y int64
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

		pos := findPosition(num)
		offset := calcOffset(pos)
		fmt.Printf("offset for %s is %d\n", line, offset)
		fmt.Print("enter a number: ")
	}

	if scanner.Err() != nil {
		os.Stderr.WriteString(fmt.Sprintf("scan error %s", scanner.Err))
	}
}

func calcOffset(pos position) int64 {
	return abs(pos.x) + abs(pos.y)
}

func findPosition(num int64) position {
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
	i := int64(1)
	currentDirection := "R"
	for true {
		for moveStep := 1; moveStep <= edgeLen; moveStep++ {
			i++
			vector := directions[currentDirection]
			pos.x += vector.x
			pos.y += vector.y
			//os.Stderr.WriteString(fmt.Sprintf("%d: step %d/%d %s (%v)\n", i, moveStep, edgeLen, currentDirection, pos))
			if i == num {
				os.Stderr.WriteString(fmt.Sprintf("%d: %v\n", num, pos))
				return pos
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
	os.Stderr.WriteString(fmt.Sprintf("------------ %v\n", pos))

	return pos
}

// Abs returns the absolute value of x.
func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
