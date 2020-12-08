package main

import (
	"fmt"

	"advent/process"
)

type day6a struct {
	data []map[rune]bool
}

func (d *day6a) Process(line string) {
	if len(line) == 0 {
		d.data = append(d.data, map[rune]bool{})
		return
	}
	curData := d.data[len(d.data)-1]
	for _, v := range line {
		curData[v] = true
	}
}

func (d *day6a) Done() {
	count := 0
	for _, v := range d.data {
		count += len(v)
	}
	fmt.Println(count)
}

type data struct {
	counts  map[rune]int
	numRows int
}

type day6b struct {
	data []data
}

func (d *day6b) Process(line string) {
	if len(line) == 0 {
		d.data = append(d.data, data{counts: map[rune]int{}})
		return
	}
	d.data[len(d.data)-1].numRows ++
	curData := d.data[len(d.data)-1].counts
	for _, v := range line {
		curData[v]++
	}
}

func (d *day6b) Done() {
	count := 0
	for _, v := range d.data {
		for _, c := range v.counts {
			if c == v.numRows {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	process.ProcessData(&day6b{data: []data{{counts: map[rune]int{}}}}, "data6.txt")
}
