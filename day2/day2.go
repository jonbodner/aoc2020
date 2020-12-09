package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	processor(&day2b{})
}

type day2a struct {
	count int
}

func (d *day2a) process(line string) {
	//1-3 a: abcde
	pos := 0
	dash := 0
	var char rune
	var min, max int
loop:
	for {
		switch line[pos] {
		case ':':
			break loop
		case '-':
			min, _ = strconv.Atoi(line[:pos])
			dash = pos
		case ' ':
			max, _ = strconv.Atoi(line[dash+1 : pos])
		default:
			char = rune(line[pos])
		}
		pos++
	}
	found := strings.Count(line[pos+1:], string(char))
	if found >= min && found <= max {
		d.count++
	}
}

func (d *day2a) done() {
	fmt.Println(d.count)
}

type day2b struct {
	count int
}

func (d *day2b) process(line string) {
	//1-3 a: abcde
	pos := 0
	dash := 0
	var char byte
	var min, max int
loop:
	for {
		switch line[pos] {
		case ':':
			break loop
		case '-':
			min, _ = strconv.Atoi(line[:pos])
			dash = pos
		case ' ':
			max, _ = strconv.Atoi(line[dash+1 : pos])
		default:
			char = line[pos]
		}
		pos++
	}
	found := 0
	if line[pos+1:][min] == char {
		found++
	}
	if line[pos+1:][max] == char {
		found++
	}
	if found == 1 {
		d.count++
	}
}

func (d *day2b) done() {
	fmt.Println(d.count)
}

type Processor interface {
	process(line string)
	done()
}

func processor(p Processor) {
	data, _ := ioutil.ReadFile("data2.txt")
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		curRow := sc.Text()
		p.process(curRow)
	}
	p.done()
}
