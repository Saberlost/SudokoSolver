//Johan was here
package main

import "testing"

type testLinepair struct {
	values []int
	answer bool
}

var testsLines = []testLinepair{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 8}, false},
	{[]int{1}, false},
	{[]int{1, 2, 3, 4, 0, 0, 0, 0, 0}, false},
	{[]int{4, 2, 3, 5, 1, 7, 9, 8, 6}, true}}

func TestValidateLine(t *testing.T) {
	for _, pair := range testsLines {
		v := validateLine(pair.values)
		if v != pair.answer {
			t.Error(
				"For", pair.values,
				"expected", pair.answer,
				"got", v)
		}

	}
}

type testLinePlacementData struct {
	values  []int
	value   int
	linepos int
	answer  bool
}

var testsLinesPlacement = []testLinePlacementData{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, 4, true},

	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 0}, 9, 9, true},
	{[]int{0}, 1, 1, true},
	{[]int{1, 2, 3, 4, 0, 0, 0, 0, 0}, 6, 5, true},
	{[]int{1, 2, 3, 4, 0, 0, 0, 0, 0}, 6, 4, false},
	{[]int{1, 2, 3, 4, 0, 0, 0, 0, 0}, 4, 4, true},
	{[]int{4, 2, 3, 5, 1, 7, 9, 8, 0}, 6, 9, true}}

func TestLinePlacement(t *testing.T) {
	for _, pair := range testsLinesPlacement {
		v := isValidPlaceMentLine(pair.values, pair.linepos-1, pair.value)
		if v != pair.answer {
			t.Error(
				"For", pair.values,
				"with pos:", pair.linepos,
				"has value:", pair.values[pair.linepos-1],
				"test value:", pair.value,
				"expected", pair.answer,
				"got", v)
		}

	}
}

type testRowpair struct {
	values [][]int
	pos    int
	answer bool
}

var testsRows = []testRowpair{
	{[][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, 1, true},
	{[][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {8}}, 1, false},
	{[][]int{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {0}, {0}}, 1, false},
	{[][]int{{5}, {1}, {3}, {4}, {6}, {9}, {7}, {8}, {2}}, 1, true},
	{[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}}, 2, true},
	{[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 8}}, 2, false},
	{[][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}}, 1, false}}

func TestValidateRow(t *testing.T) {
	for _, pair := range testsRows {
		v := validateRow(pair.values, pair.pos-1)
		if v != pair.answer {
			t.Error(
				"For", pair.values,
				"with pos: ", pair.pos,
				"expected", pair.answer,
				"got", v)
		}

	}
}

type testRowPlacement struct {
	values [][]int
	pos    int
	row    int
	value  int
	answer bool
}

var testsRowsPlacement = []testRowPlacement{
	{[][]int{
		{1},
		{2},
		{3},
		{4},
		{0},
		{6},
		{7},
		{8},
		{9}}, 1, 5, 5, true},

	{[][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
		{0},
		{0},
		{0},
		{8}}, 1, 5, 5, false},
	{[][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
		{0},
		{0},
		{0},
		{8}}, 1, 5, 6, false},

	{[][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{0, 0},
		{0, 7},
		{0, 8},
		{0, 9}}, 2, 6, 6, true},
	{[][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{0, 0},
		{0, 7},
		{0, 8},
		{0, 9}}, 2, 6, 5, false}}

func TestRowPlacment(t *testing.T) {
	for _, pair := range testsRowsPlacement {
		v := isValidPlacementRow(pair.values, pair.pos-1, pair.row-1, pair.value)
		if v != pair.answer {
			t.Error(
				"When validate row placement for", pair.values,
				"with row: ", pair.row,
				"and pos : ", pair.pos,
				//"has value:", pair.values[pair.pos-1][pair.row-1],
				"test value:", pair.value,
				"expected", pair.answer,
				"got", v)
		}

	}
}

type testBoxpairs struct {
	values [][]int
	square int
	answer bool
}

var testsBoxs = []testBoxpairs{
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}, 0, true},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 8}}, 0, false},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 0}}, 0, false},
	{[][]int{
		{5, 1, 3},
		{4, 6, 9},
		{7, 8, 2}}, 0, true},
	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 0, false},
	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 1, true},
	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 4, 6},
		{0, 0, 0, 7, 8, 9}}, 1, false},
	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9},
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 4, true},
	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9},
		{0, 0, 0, 1, 8, 3},
		{0, 0, 0, 4, 8, 6},
		{0, 0, 0, 7, 8, 9}}, 4, false},
	{[][]int{
		{4, 5, 6, 1, 8, 3},
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 7, 2, 9},
		{7, 8, 9, 1, 8, 3},
		{1, 2, 3, 4, 8, 6},
		{4, 5, 6, 7, 8, 9}}, 4, false},
	{[][]int{
		{4, 5, 6, 1, 8, 3},
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 7, 2, 9},
		{7, 8, 9, 1, 8, 3},
		{1, 2, 3, 4, 2, 6},
		{4, 5, 6, 7, 5, 9}}, 4, true}}

func TestBox(t *testing.T) {
	for _, pair := range testsBoxs {
		v := validateBox(pair.values, pair.square)
		if v != pair.answer {
			t.Error(
				"For", pair.values,
				"with index: ", pair.square,
				"expected", pair.answer,
				"got", v)
		}

	}
}

var testsBoxsPlacment = []testRowPlacement{
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}, 2, 2, 4, false},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 8}}, 2, 3, 9, true},

	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 1, 1, 2, true},
	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 4, 1, 1, false},
	{[][]int{
		{0, 0, 0, 0, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 4, 1, 1, true},
	{[][]int{
		{0, 0, 0, 0, 2, 3},
		{0, 0, 0, 4, 0, 6},
		{0, 0, 0, 7, 8, 9}}, 5, 2, 1, true},

	{[][]int{
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9},
		{0, 0, 0, 1, 2, 3},
		{0, 0, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9}}, 2, 2, 2, true},
	{[][]int{
		{0, 2, 0, 1, 2, 3},
		{0, 1, 0, 4, 5, 6},
		{0, 0, 0, 7, 8, 9},
		{0, 2, 0, 1, 0, 3},
		{0, 1, 0, 4, 0, 6},
		{0, 0, 0, 7, 8, 9}}, 5, 5, 2, true},
	{[][]int{
		{4, 5, 6, 1, 8, 3},
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 7, 2, 0},
		{7, 8, 9, 1, 8, 3},
		{1, 2, 3, 4, 8, 6},
		{4, 5, 6, 7, 8, 9}}, 6, 3, 9, true}}

func TestBoxPlacement(t *testing.T) {
	for _, pair := range testsBoxsPlacment {
		v := IsValidPlaceMentBox(pair.values, pair.pos-1, pair.row-1, pair.value)
		if v != pair.answer {
			t.Error(
				"When validate  placement in box for", pair.values,
				"with row: ", pair.row,
				"and pos : ", pair.pos,
				//"has value:", pair.values[pair.pos-1][pair.row-1],
				"test value:", pair.value,
				"expected", pair.answer,
				"got", v)
		}

	}
}

func TestGetSquareFromCords(t *testing.T) {
	if getSquareFromCords(0, 0) != 0 {
		t.Error("wrong square")
	}
	if getSquareFromCords(1, 1) != 0 {
		t.Error("wrong square")
	}
	if getSquareFromCords(3, 3) != 4 {
		t.Error("wrong square")
	}
	if getSquareFromCords(8, 8) != 8 {
		t.Error("wrong square")
	}
}
