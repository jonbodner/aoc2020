package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"advent/process"
)

type day5b struct {
	data []string
}

func (d *day5b) Process(line string) {
	d.data = append(d.data, line)
}

func (d *day5b) Done() {
	ids := make([]int, 0, len(d.data))
	for _, v := range d.data {
		fmt.Println(v)
		row, col, id := calcRCID2(v)
		fmt.Println(row, col, id)
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for i, v := range ids[:len(ids)-1] {
		if ids[i+1]-v > 1 {
			fmt.Println(v, ids[i+1])
		}
	}
}

type day5a struct {
	data []string
}

func (d *day5a) Process(line string) {
	d.data = append(d.data, line)
}

/*
For example, consider just the first seven characters of FBFBBFFRLR:

Start by considering the whole range, rows 0 through 127.
F means to take the lower half, keeping rows 0 through 63.
B means to take the upper half, keeping rows 32 through 63.
F means to take the lower half, keeping rows 32 through 47.
B means to take the upper half, keeping rows 40 through 47.
B keeps rows 44 through 47.
F keeps rows 44 through 45.
The final F keeps the lower of the two, row 44.
The last three characters will be either L or R; these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7). The same process as above proceeds again, this time with only three steps. L means to keep the lower half, while R means to keep the upper half.

For example, consider just the last 3 characters of FBFBBFFRLR:

Start by considering the whole range, columns 0 through 7.
R means to take the upper half, keeping columns 4 through 7.
L means to take the lower half, keeping columns 4 through 5.
The final R keeps the upper of the two, column 5.
So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.

Every seat also has a unique seat ID: multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.

Here are some other boarding passes:

BFFFBBFRRR: row 70, column 7, seat ID 567.
FFFBBBFRRR: row 14, column 7, seat ID 119.
BBFFBBFRLL: row 102, column 4, seat ID 820.
*/
func (d *day5a) Done() {
	maxID := 0
	for _, v := range d.data {
		fmt.Println(v)
		row, col, id := calcRCID(v)
		fmt.Println(row, col, id)
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println(maxID)
}

func calcRCID(v string) (int, int, int) {
	minR := 0
	maxR := 127
	for _, c := range v[:7] {
		switch c {
		case 'F':
			maxR = (minR + maxR) / 2
		case 'B':
			minR = (maxR + minR) / 2
		}
	}
	minC := 0
	maxC := 7
	for _, c := range v[7:] {
		switch c {
		case 'L':
			maxC = (minC + maxC) / 2
		case 'R':
			minC = (maxC + minC) / 2
		}
	}
	id := maxR*8 + maxC
	return maxR, maxC, id
}

func calcRCID2(v string) (int, int, int) {
	v = strings.ReplaceAll(v, "F", "0")
	v = strings.ReplaceAll(v, "B", "1")
	v = strings.ReplaceAll(v, "L", "0")
	v = strings.ReplaceAll(v, "R", "1")
	row, err := strconv.ParseInt(v[:7], 2, 64)
	if err != nil {
		panic(err)
	}
	col, err := strconv.ParseInt(v[7:], 2, 64)
	if err != nil {
		panic(err)
	}
	id := row*8 + col
	return int(row), int(col), int(id)
}

func main() {
	process.ProcessData(&day5b{}, "data5.txt")
}
