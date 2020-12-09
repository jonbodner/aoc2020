package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"advent/process"
)

type instruction interface {
	run(ac int, pc int) (acOut int, pcOut int)
}

type nop int

func (n nop) run(ac int, pc int) (int, int) {
	return ac, pc + 1
}

type jmp int

func (j jmp) run(ac int, pc int) (int, int) {
	return ac, pc + int(j)
}

type acc int

func (a acc) run(ac int, pc int) (int, int) {
	return ac + int(a), pc + 1
}

type day8a struct {
	instructions []instruction
}

func (d *day8a) Process(line string) {
	parts := strings.Split(line, " ")
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	switch parts[0] {
	case "nop":
		d.instructions = append(d.instructions, nop(i))
	case "jmp":
		d.instructions = append(d.instructions, jmp(i))
	case "acc":
		d.instructions = append(d.instructions, acc(i))
	}
}

func (d *day8a) Done() {
	ac := 0
	pc := 0
	seen := map[int]int{}
	for seen[pc] != 2 {
		v := d.instructions[pc]
		ac, pc = v.run(ac, pc)
		seen[pc]++
		switch v.(type) {
		case nop:
			fmt.Println("nop ", v)
		case acc:
			fmt.Println("acc ", v)
		case jmp:
			fmt.Println("jmp", v)
		}
	}
	fmt.Println(ac, pc)
}

type day8b struct {
	instructions []instruction
}

func (d *day8b) Process(line string) {
	parts := strings.Split(line, " ")
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	switch parts[0] {
	case "nop":
		d.instructions = append(d.instructions, nop(i))
	case "jmp":
		d.instructions = append(d.instructions, jmp(i))
	case "acc":
		d.instructions = append(d.instructions, acc(i))
	}
}

func doIt(instructions []instruction, flipped int) {
	ac := 0
	pc := 0
	seen := map[int]int{}
	for pc < len(instructions) && seen[pc] != 2 {
		v := instructions[pc]
		ac, pc = v.run(ac, pc)
		seen[pc]++
	}
	if pc >= len(instructions) {
		fmt.Println("success!", ac, pc, flipped, instructions[flipped])
		return
	}
}

func (d *day8b) Done() {
	var wg sync.WaitGroup
	tried := 0
	for i, v := range d.instructions {
		switch v := v.(type) {
		case nop:
			tried++
			wg.Add(1)
			inst := make([]instruction, len(d.instructions))
			copy(inst, d.instructions)
			i := i
			inst[i] = jmp(v)
			go func() {
				defer wg.Done()
				doIt(inst, i)
			}()
		case jmp:
			tried++
			wg.Add(1)
			inst := make([]instruction, len(d.instructions))
			copy(inst, d.instructions)
			i := i
			inst[i] = nop(v)
			go func() {
				defer wg.Done()
				doIt(inst, i)
			}()
		}
	}
	wg.Wait()
	fmt.Println(tried)
}

func main() {
	process.ProcessData(&day8b{}, "data8.txt")
}
