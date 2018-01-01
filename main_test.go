package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPositions(t *testing.T) {
	Convey("Given a known-position number", t, func() {
		expected := []position{
			position{0, 0},
			position{1, 0},
			position{1, -1},
			position{0, -1},
			position{-1, -1},
			position{-1, 0},
			position{-1, 1},
			position{0, 1},
			position{1, 1},
			position{2, 1},
		}
		for i, position := range expected {
			checkPosition(t, int64(i+1), position)
		}

	})
}

type offsetTest struct {
	position position
	offset   int
}

func TestOffsets(t *testing.T) {
	Convey("Given a coordinate", t, func() {
		expected := []offsetTest{offsetTest{position{0, 0}, 0},
			offsetTest{position{1, 0}, 1},
			offsetTest{position{1, -1}, 2},
			offsetTest{position{-100, 0}, 100},
			offsetTest{position{-100, 99}, 199},
		}
		for _, offsetTest := range expected {
			pos := offsetTest.position
			expectedOffset := offsetTest.offset
			Convey(fmt.Sprintf("offset correct for %d,%d", pos.x, pos.y), func() {
				offset := calcOffset(pos)
				So(offset, ShouldEqual, expectedOffset)
			})
		}
	})
}

func checkPosition(t *testing.T, num int64, expectedpos position) {
	pos := findPosition(int(num))
	Convey(fmt.Sprintf("coordinates are correct for %d", num), func() {
		So(pos, ShouldResemble, expectedpos)
	})
}

func TestExceedingValues(t *testing.T) {
	Convey("Given a number", t, func() {
		nextValue := spiralEntryExceeding(5)
		So(nextValue, ShouldEqual, 10)
		expected := []int{2, 4, 4, 5, 10}
		for i, nextExpected := range expected {
			Convey(fmt.Sprintf("the next seen value in the spiral after %d is %d", i+1, nextExpected), func() {
				nextValue := spiralEntryExceeding(i + 1)
				So(nextValue, ShouldEqual, nextExpected)
			})
		}
	})
}

func TestSumNeighbours(t *testing.T) {
	Convey("third iteration of spiral", t, func() {
		matrix := [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		}
		pos := position{2, 1}
		cellValue := sumNeighbours(matrix, pos)
		So(cellValue, ShouldEqual, 2)
	})
	Convey("fourth iteration", t, func() {
		matrix := [][]int{
			{0, 0, 0, 0},
			{0, 0, 2, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		}
		pos := position{1, 1}
		cellValue := sumNeighbours(matrix, pos)
		So(cellValue, ShouldEqual, 4)
	})
	Convey("fifth iteration", t, func() {
		matrix := [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 4, 2, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		pos := position{1, 1}
		cellValue := sumNeighbours(matrix, pos)
		So(cellValue, ShouldEqual, 5)
	})
}
