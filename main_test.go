package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPositions(t *testing.T) {
	Convey("Given a known-position number", t, func() {
		expected := []position{position{0, 0}, position{1, 0}, position{1, -1}}
		for i, position := range expected {
			checkPosition(t, int64(i+1), position)
		}

		checkPosition(t, 5, position{-1, -1})
		checkPosition(t, 6, position{-1, 0})
		checkPosition(t, 7, position{-1, 1})
		checkPosition(t, 8, position{0, 1})
		checkPosition(t, 9, position{1, 1})
		checkPosition(t, 10, position{2, 1})
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
	pos := findPosition(num)
	Convey(fmt.Sprintf("coordinates are correct for %d", num), func() {
		So(pos, ShouldResemble, expectedpos)
	})
}
