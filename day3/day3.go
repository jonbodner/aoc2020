package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	processor(&day3b{})

}

type day3a struct {
	m []string
}

func (d *day3a) process(line string) {
	d.m = append(d.m, line)
}

func (d *day3a) done() {
	x := 0
	y := 0
	count := 0
	for y < len(d.m)-1 {
		x += 3
		y += 1
		if x >= len(d.m[y]) {
			x = x % len(d.m[y])
		}
		if d.m[y][x] == '#' {
			count++
		}
	}
	fmt.Println(count)
}

type day3b struct {
	m []string
}

func (d *day3b) process(line string) {
	d.m = append(d.m, line)
}

/*
Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
*/
func (d *day3b) done() {
	data := []struct {
		incX int
		incY int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	mults := 1
	for _, ds := range data {
		x := 0
		y := 0
		count := 0
		for y < len(d.m)-1 {
			x += ds.incX
			y += ds.incY
			if x >= len(d.m[y]) {
				x = x % len(d.m[y])
			}
			if d.m[y][x] == '#' {
				count++
			}
		}
		fmt.Println(ds, count)
		mults *= count
	}
	fmt.Println(mults)
}

type Processor interface {
	process(line string)
	done()
}

func processor(p Processor) {
	data, _ := ioutil.ReadFile("data3.txt")
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		curRow := sc.Text()
		p.process(curRow)
	}
	p.done()
}
