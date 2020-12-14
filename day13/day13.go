package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	ProcessDataA(&day13a{}, "data13.txt")
	ProcessDataB(&day13b{}, "data13.txt")
}

type day13a struct {
	arrival int
	busses []int
}

func (a *day13a) Calculate() {
	curID := -1
	minTime := math.MaxInt64
	for _, v := range a.busses {
		mults := a.arrival / v
		if a.arrival == v * mults {
			// exact!
			curID = v
			minTime = 0
			break
		}
		diff := v * (mults+1) - a.arrival
		if diff < minTime {
			minTime = diff
			curID = v
		}
	}
	fmt.Println(curID, minTime, curID * minTime)
}

type day13b struct {
	arrival int
	busses []int
	biggest int
	biggestPos int
}

func (a *day13b) Calculate() {
	for i := 1;true;i++ {
		val := a.biggest * i
		start := val - a.biggestPos
		success := true
		for j, v := range a.busses {
			if a.busses[j] == -1 {
				continue
			}
			if (start + j) % v != 0 {
				success = false
				break
			}
		}
		if success {
			fmt.Println(start)
			return
		}
	}
}

func (a *day13b) CalculateBroken() {
	factor := 1
	for _, v := range a.busses {
		if v == -1 {
			continue
		}
		factor *= v
	}
	fmt.Println(factor)
	for i := 1;true;i++ {
		if i % 1000 == 0 {
			fmt.Println(i)
		}
		val := i * factor
		start := val - a.biggestPos
		success := true
		for j, v := range a.busses {
			if a.busses[j] == -1 {
				continue
			}
			if (start + j) % v != 0 {
				success = false
				break
			}
		}
		if success {
			fmt.Println(start)
			return
		}
	}
}

func ProcessDataA(p *day13a, filename string) {
	data, _ := ioutil.ReadFile(filename)
	sc := bufio.NewScanner(bytes.NewReader(data))
	// arrival timestamp
	sc.Scan()
	curRow := sc.Text()
	var err error
	p.arrival, err = strconv.Atoi(curRow)
	if err != nil {
		panic(err)
	}
	// busses
	sc.Scan()
	curRow = sc.Text()
	parts := strings.Split(curRow,",")
	for _, v := range parts {
		if v == "x" {
			continue
		}
		id, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		p.busses = append(p.busses, id)
	}
	p.Calculate()
}

func ProcessDataB(p *day13b, filename string) {
	data, _ := ioutil.ReadFile(filename)
	sc := bufio.NewScanner(bytes.NewReader(data))
	// arrival timestamp -- skip
	sc.Scan()
	// busses
	sc.Scan()
	curRow := sc.Text()
	parts := strings.Split(curRow,",")
	for _, v := range parts {
		if v == "x" {
			p.busses = append(p.busses,-1)
			continue
		}
		id, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if id > p.biggest {
			p.biggest = id
			p.biggestPos = len(p.busses)
		}
		p.busses = append(p.busses, id)
	}
	p.Calculate()
}
