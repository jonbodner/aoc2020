package main

import (
	"fmt"
	"sort"
	"strconv"

	"advent/process"
)

type day10a struct {
	nums []int
}

func (d *day10a) Process(line string) {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	d.nums = append(d.nums, num)
}

func (d *day10a) Done() {
	sort.Ints(d.nums)
	num1 := 0
	num3 := 0
	curVal := 0
	for _, v := range d.nums {
		if v-curVal == 1 {
			num1++
		}
		if v-curVal == 3 {
			num3++
		}
		curVal = v
	}
	num3++
	fmt.Println(num1, num3, num1*num3)
}

type day10b struct {
	nums []int
}

func (d *day10b) Process(line string) {
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	d.nums = append(d.nums, num)
}

func (d *day10b) Done() {
	sort.Ints(d.nums)
	lastNum := d.nums[len(d.nums)-1] + 3
	vals := append([]int{}, 0)
	vals = append(vals, d.nums...)
	vals = append(vals, lastNum)
	found := make([]int, len(vals))
	count := findAllChains(vals, found)
	fmt.Println(count)
}

func findAllChains(nums []int, found []int) int {
	if found[0] != 0 {
		return found[0]
	}
	curVal := nums[0]
	i := 0
	for ; i < len(nums)-2 && nums[i+2]-curVal > 3; i++ {
		curVal = nums[i+1]
	}
	if i >= len(nums)-2 {
		found[i] = 1
		return 1
	}
	total := 0
	for j := i + 1; j < len(nums) && nums[j]-nums[i] <= 3; j++ {
		total += findAllChains(nums[j:], found[j:])
	}
	found[i] = total
	return total
}

func main() {
	process.ProcessData(&day10b{}, "data10.txt")
}
