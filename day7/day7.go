package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent/process"
)

type result struct {
	count int
	color string
}

type day7a struct {
	rules map[string][]result
}

/*
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
*/
func (d *day7a) Process(line string) {
	endPos1 := strings.Index(line, "bags")
	color := line[:endPos1-1]
	var results []result
	allResults := strings.Split(line[endPos1+13:], ", ")
	for _, v := range allResults {
		if v == "no other bags." {
			break
		}
		words := strings.Split(v, " ")
		//number color color
		count, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}
		results = append(results, result{count: count, color: words[1] + " " + words[2]})
	}
	d.rules[color] = results
}

/*
You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors
would be valid for the outermost bag? (In other words: how many colors can, eventually,
contain at least one shiny gold bag?)

4
*/
func (d *day7a) Done() {
	curColors := map[string]bool{"shiny gold": true}
	seenColors := map[string]bool{}
	for len(curColors) > 0 {
		newColors := map[string]bool{}
		for color, results := range d.rules {
			for _, v := range results {
				if _, ok := curColors[v.color]; ok {
					if _, ok := seenColors[color]; !ok {
						newColors[color] = true
					}
				}
			}
		}
		for k := range curColors {
			seenColors[k] = true
		}
		curColors = newColors
	}
	// includes shiny gold, which doesn't count
	fmt.Println(len(seenColors) - 1)
}

type day7b struct {
	rules map[string][]result
}

/*
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
*/
func (d *day7b) Process(line string) {
	endPos1 := strings.Index(line, "bags")
	color := line[:endPos1-1]
	var results []result
	allResults := strings.Split(line[endPos1+13:], ", ")
	for _, v := range allResults {
		if v == "no other bags." {
			break
		}
		words := strings.Split(v, " ")
		//number color color
		count, err := strconv.Atoi(words[0])
		if err != nil {
			panic(err)
		}
		results = append(results, result{count: count, color: words[1] + " " + words[2]})
	}
	d.rules[color] = results
}

/*
How many individual bags are required inside your single shiny gold bag?
*/
func (d *day7b) Done() {
	seen := map[string]int{}
	count := d.doneInner("shiny gold", seen)
	fmt.Println(count)
}

func (d *day7b) doneInner(curColor string, seen map[string]int) int {
	count := 0
	curResults := d.rules[curColor]
	for _, result := range curResults {
		if num, ok := seen[result.color]; ok {
			fmt.Println(result.color, num)
			count += result.count + (result.count * num)
		} else {
			count += result.count + (result.count * d.doneInner(result.color, seen))
		}
	}
	seen[curColor] = count
	fmt.Println(curColor, count)
	return count
}

func main() {
	process.ProcessData(&day7b{rules: map[string][]result{}}, "data7.txt")
}
