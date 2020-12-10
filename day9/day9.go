package main

import (
	"fmt"
	"math"
	"strconv"

	"advent/process"
)

type day9b struct {
	preamble int
	nums     []int
}

func (d *day9b) Process(line string) {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	d.nums = append(d.nums, num)
}

func (d *day9b) Done() {
	cur := d.preamble
	curNum := d.nums[cur]
	for findPair(curNum, d.nums[cur-d.preamble:cur]) {
		cur++
		curNum = d.nums[cur]
	}
	fmt.Println(curNum)
	vals := findNums(curNum, d.nums[:cur])
	mx := max(vals)
	mn := min(vals)
	fmt.Println(mn, mx, mn+mx)
}

func min(ints []int) int {
	out := math.MaxInt64
	for _, v := range ints {
		if v < out {
			out = v
		}
	}
	return out
}

func max(ints []int) int {
	out := 0
	for _, v := range ints {
		if v > out {
			out = v
		}
	}
	return out
}

func findNums(sum int, ints []int) []int {
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			curSum := 0
			for k := i; k < j; k++ {
				curSum += ints[k]
				if curSum == sum {
					return ints[i:j]
				}
				if curSum > sum {
					break
				}
			}
		}
	}
	return nil
}

type day9a struct {
	preamble int
	nums     []int
}

func (d *day9a) Process(line string) {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	d.nums = append(d.nums, num)
}

func (d *day9a) Done() {
	cur := d.preamble
	for findPair(d.nums[cur], d.nums[cur-d.preamble:cur]) {
		cur++
	}
	fmt.Println(d.nums[cur])
}

func findPair(sum int, ints []int) bool {
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i]+ints[j] == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	process.ProcessData(&day9b{preamble: 25}, "data9.txt")
}
