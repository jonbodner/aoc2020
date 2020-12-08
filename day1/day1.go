package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

/*
Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
 */
func main() {
	processor(&holder2{})
}

type holder []int

func (h *holder) process(n int) {
	*h = append(*h, n)
}

func (h *holder) done() {
	for i := 0;i<len(*h)-1;i++ {
		for j := i +1;j< len(*h);j++ {
			if (*h)[i] + (*h)[j] == 2020 {
				fmt.Println((*h)[i] * (*h)[j])
				return
			}
		}
	}
}

type holder2 []int

func (h *holder2) process(n int) {
	*h = append(*h, n)
}

func (h *holder2) done() {
	for i := 0;i<len(*h)-2;i++ {
		for j := i +1;j< len(*h)-1;j++ {
			for k := j + 1; k<len(*h);k++ {
				if (*h)[i] + (*h)[j] + (*h)[k]== 2020 {
					fmt.Println((*h)[i] * (*h)[j] * (*h)[k])
					return
				}
			}
		}
	}
}

type Processor interface {
	process(n int)
	done()
}

func processor(p Processor) {
	data, _ := ioutil.ReadFile("data.txt")
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		curRow := sc.Text()
		num, _ := strconv.Atoi(curRow)
		p.process(num)
	}
	p.done()
}
