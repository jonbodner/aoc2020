package main

import (
	"fmt"
	"math"
	"strconv"

	"advent/process"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

type day12a struct {
	posX   int // lower is west, higher is east
	posY   int // lower is north, higher is south
	facing direction
}

func (d *day12a) Process(line string) {
	num, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	switch line[0] {
	case 'F':
		switch d.facing {
		case north:
			d.posY -= num
		case south:
			d.posY += num
		case east:
			d.posX += num
		case west:
			d.posX -= num
		default:
			panic("bad position")
		}
	case 'N':
		d.posY -= num
	case 'S':
		d.posY += num
	case 'E':
		d.posX += num
	case 'W':
		d.posX -= num
	case 'L':
		fmt.Println("left before:", d.facing, num)
		amount := num / 90
		d.facing -= direction(amount)
		if d.facing < 0 {
			d.facing = 4 + d.facing
		}
		d.facing = d.facing % 4
		fmt.Println("left after:", d.facing)
	case 'R':
		fmt.Println("right before:", d.facing, num)
		amount := num / 90
		d.facing += direction(amount)
		d.facing = d.facing % 4
		fmt.Println("right after:", d.facing)
	default:
		panic("bad key")
	}
}

func (d *day12a) Done() {
	fmt.Println(int(math.Abs(float64(d.posX))) + int(math.Abs(float64(d.posY))))
}

type day12b struct {
	posX   int // lower is west, higher is east
	posY   int // lower is north, higher is south
	facing direction

	wpX int
	wpY int
}

/*
Almost all of the actions indicate how to move a waypoint which is relative to the ship's position:

Action N means to move the waypoint north by the given value.
Action S means to move the waypoint south by the given value.
Action E means to move the waypoint east by the given value.
Action W means to move the waypoint west by the given value.
Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
Action F means to move forward to the waypoint a number of times equal to the given value.
The waypoint starts 10 units east and 1 unit north relative to the ship. The waypoint is relative to the ship; that is, if the ship moves, the waypoint moves with it.
*/
func (d *day12b) Process(line string) {
	num, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	switch line[0] {
	case 'F':
		for i := 0; i < num; i++ {
			d.posX += d.wpX
			d.posY += d.wpY
		}
	case 'N':
		d.wpY -= num
	case 'S':
		d.wpY += num
	case 'E':
		d.wpX += num
	case 'W':
		d.wpX -= num
	case 'R':
		switch num {
		case 90:
			tmp := d.wpX
			d.wpX = -d.wpY
			d.wpY = tmp
		case 180:
			d.wpX = -d.wpX
			d.wpY = -d.wpY
		case 270:
			tmp := d.wpX
			d.wpX = d.wpY
			d.wpY = -tmp
		}
	case 'L':
		switch num {
		case 270:
			tmp := d.wpX
			d.wpX = -d.wpY
			d.wpY = tmp
		case 180:
			d.wpX = -d.wpX
			d.wpY = -d.wpY
		case 90:
			tmp := d.wpX
			d.wpX = d.wpY
			d.wpY = -tmp
		}
	default:
		panic("bad key")
	}
	fmt.Println(d)
}

func (d *day12b) Done() {
	fmt.Println(int(math.Abs(float64(d.posX))) + int(math.Abs(float64(d.posY))))
}

func main() {
	process.ProcessData(&day12b{wpX: 10, wpY: -1, facing: east}, "data12.txt")
}
