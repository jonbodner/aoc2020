package main

import "testing"

func TestIt(t *testing.T) {
	data := []struct {
		in  string
		row int
		col int
		id  int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}
	for _, v := range data {
		t.Run(v.in, func(t *testing.T) {
			row, col, id := calcRCID(v.in)
			if row != v.row || col != v.col || id != v.id {
				t.Error(row, col, id, v)
			}
		})
	}
}

func TestIt2(t *testing.T) {
	data := []struct {
		in  string
		row int
		col int
		id  int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}
	for _, v := range data {
		t.Run(v.in, func(t *testing.T) {
			row, col, id := calcRCID2(v.in)
			if row != v.row || col != v.col || id != v.id {
				t.Error(row, col, id, v)
			}
		})
	}
}
