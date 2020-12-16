package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent/process"
)

type day14a struct {
	memory  map[int]uint64
	orMask  uint64 // has 1s
	andMask uint64 // has 0s
}

/*
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
*/
func (d *day14a) Process(line string) {
	if strings.HasPrefix(line, "mask = ") {
		d.orMask = 0
		d.andMask = 1<<36 - 1
		fmt.Printf("and: %36b\n", d.andMask)
		fmt.Printf("or: %36b\n", d.orMask)
		for i, c := range line[7:] {
			switch c {
			case 'X':
				continue
			case '1':
				d.orMask = d.orMask | uint64(1<<(35-i))
			case '0':
				d.andMask = d.andMask ^ uint64(1<<(35-i))
			default:
				panic("unknown: " + line)
			}
		}
		fmt.Printf("and: %36b\n", d.andMask)
		fmt.Printf("or: %36b\n", d.orMask)
	}
	if strings.HasPrefix(line, "mem[") {
		end := strings.Index(line, "]")
		offset, err := strconv.Atoi(line[4:end])
		if err != nil {
			panic(err)
		}
		num, err := strconv.ParseUint(line[end+4:], 10, 64)
		if err != nil {
			panic(err)
		}
		realNum := (num & d.andMask) | d.orMask
		d.memory[offset] = realNum
	}
}

func (d *day14a) Done() {
	var sum uint64
	for _, v := range d.memory {
		sum += v
	}
	fmt.Println(sum)
}

type day14b struct {
	memory    map[uint64]uint64
	orMask    uint64 // has 1s
	andMask   uint64 // has 0s
	floatBits []int
}

/*
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
*/
func (d *day14b) Process(line string) {
	if strings.HasPrefix(line, "mask = ") {
		d.orMask = 0
		d.andMask = 1<<36 - 1
		d.floatBits = nil
		fmt.Println("msk:", line[7:])
		for i, c := range line[7:] {
			switch c {
			case 'X':
				d.floatBits = append(d.floatBits, 35-i)
			case '1':
				d.orMask = d.orMask | uint64(1<<(35-i))
			case '0':
				//d.andMask = d.andMask ^ uint64(1<<(35-i))
			default:
				panic("unknown: " + line)
			}
		}
		fmt.Printf("and: %036b\n", d.andMask)
		fmt.Printf("or : %036b\n", d.orMask)
		fmt.Printf("floats: %v\n", d.floatBits)
	}
	if strings.HasPrefix(line, "mem[") {
		end := strings.Index(line, "]")
		offset, err := strconv.ParseUint(line[4:end], 10, 64)
		if err != nil {
			panic(err)
		}
		num, err := strconv.ParseUint(line[end+4:], 10, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ori: %036b\n", offset)
		offset = (offset & d.andMask) | d.orMask

		max := 1 << len(d.floatBits)
		for i := 0; i < max; i++ {
			orMask2, andMask2 := calcMasks(i, d.floatBits)
			subOffset := (offset & andMask2) | orMask2
			fmt.Println("writing", num, "to", subOffset)
			d.memory[subOffset] = num
		}
	}
}

func calcMasks(curInt int, bits []int) (uint64, uint64) {
	var orMask uint64 = 0
	var andMask uint64 = 1<<36 - 1
	for i, v := range bits {
		if curInt&(1<<(len(bits)-i-1)) != 0 {
			orMask |= uint64(1 << v)
		} else {
			andMask ^= uint64(1 << v)
		}
	}
	fmt.Printf("and: %036b\n", andMask)
	fmt.Printf("or : %036b\n", orMask)
	return orMask, andMask
}

func (d *day14b) Done() {
	var sum uint64
	for _, v := range d.memory {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	process.ProcessData(&day14b{memory: map[uint64]uint64{}}, "data14.txt")
}
