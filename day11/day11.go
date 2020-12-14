package main

import (
	"bytes"
	"fmt"

	"advent/process"
)

/*
The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
*/

type state int

func (s state) String() string {
	switch s {
	case floor:
		return "."
	case empty:
		return "L"
	case occupied:
		return "#"
	}
	panic("illegal state")
}

func toState(b rune) state {
	switch b {
	case '.':
		return floor
	case 'L':
		return empty
	case '#':
		return occupied
	}
	panic("illegal state")
}

const (
	floor state = iota
	empty
	occupied
)

type day11a struct {
	layout [][]state
}

func (d *day11a) Process(line string) {
	var curRow []state
	for _, v := range line {
		curRow = append(curRow, toState(v))
	}
	d.layout = append(d.layout, curRow)
}

/*
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
Floor (.) never changes; seats don't move, and nobody sits on the floor.
*/
func (d *day11a) Done() {
	curBoard := d.layout
	fmt.Println(print(curBoard))
	fmt.Println()
	for {
		newBoard := copyBoard(curBoard)
		for k, v := range curBoard {
			for k2, v2 := range v {
				switch v2 {
				case floor:
					continue
				case empty:
					if countNeighbors(k, k2, curBoard) == 0 {
						newBoard[k][k2] = occupied
					}
				case occupied:
					if countNeighbors(k, k2, curBoard) >= 4 {
						newBoard[k][k2] = empty
					}
				}
			}
		}
		if equal(curBoard, newBoard) {
			break
		}
		curBoard = newBoard
		fmt.Println(print(curBoard))
		fmt.Println()
	}
	fmt.Println(countOccupied(curBoard))
}

type day11b struct {
	layout [][]state
}

func (d *day11b) Process(line string) {
	var curRow []state
	for _, v := range line {
		curRow = append(curRow, toState(v))
	}
	d.layout = append(d.layout, curRow)
}

/*
If a seat is empty (L) and there are no occupied seats VISIBLE to it, the seat becomes occupied.
If a seat is occupied (#) and FIVE or more seats VISIBLE to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
Floor (.) never changes; seats don't move, and nobody sits on the floor.
*/
func (d *day11b) Done() {
	curBoard := d.layout
	fmt.Println(print(curBoard))
	fmt.Println()
	for {
		newBoard := copyBoard(curBoard)
		for k, v := range curBoard {
			for k2, v2 := range v {
				switch v2 {
				case floor:
					continue
				case empty:
					if countVisible(k, k2, curBoard) == 0 {
						newBoard[k][k2] = occupied
					}
				case occupied:
					if countVisible(k, k2, curBoard) >= 5 {
						newBoard[k][k2] = empty
					}
				}
			}
		}
		if equal(curBoard, newBoard) {
			break
		}
		curBoard = newBoard
		fmt.Println(print(curBoard))
		fmt.Println()
	}
	fmt.Println(countOccupied(curBoard))
}

func print(board [][]state) string {
	var out bytes.Buffer
	for _, v := range board {
		for _, v2 := range v {
			out.WriteString(v2.String())
		}
		out.WriteByte('\n')
	}
	return out.String()
}

func countNeighbors(x, y int, board [][]state) int {
	var count int
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			curX := x + i
			curY := y + j
			if curX >= 0 && curY >= 0 && curX < len(board) && curY < len(board[0]) && board[curX][curY] == occupied {
				count++
			}
		}
	}
	return count
}

func countVisible(x, y int, board [][]state) int {
	var count int
	// constant X to 0 (sub)
	for i := x - 1; i >= 0; i-- {
		if board[i][y] == empty {
			break
		}
		if board[i][y] == occupied {
			count++
			break
		}
	}
	// constant X to board width (add)
	for i := x + 1; i < len(board); i++ {
		if board[i][y] == empty {
			break
		}
		if board[i][y] == occupied {
			count++
			break
		}
	}
	// constant Y to 0 (sub)
	for j := y - 1; j >= 0; j-- {
		if board[x][j] == empty {
			break
		}
		if board[x][j] == occupied {
			count++
			break
		}
	}
	// constant Y to board height (add)
	for j := y + 1; j < len(board[0]); j++ {
		if board[x][j] == empty {
			break
		}
		if board[x][j] == occupied {
			count++
			break
		}
	}

	// sub X, sub Y
	for v := 1; x-v >= 0 && y-v >= 0; v++ {
		if board[x-v][y-v] == empty {
			break
		}
		if board[x-v][y-v] == occupied {
			count++
			break
		}
	}
	// sub X, add Y
	for v := 1; x-v >= 0 && y+v < len(board[0]); v++ {
		if board[x-v][y+v] == empty {
			break
		}
		if board[x-v][y+v] == occupied {
			count++
			break
		}
	}
	// add X, sub Y
	for v := 1; x+v < len(board) && y-v >= 0; v++ {
		if board[x+v][y-v] == empty {
			break
		}
		if board[x+v][y-v] == occupied {
			count++
			break
		}
	}
	// add X, add Y
	for v := 1; x+v < len(board) && y+v < len(board[0]); v++ {
		if board[x+v][y+v] == empty {
			break
		}
		if board[x+v][y+v] == occupied {
			count++
			break
		}
	}
	return count
}

func countOccupied(board [][]state) int {
	count := 0
	for _, v := range board {
		for _, v2 := range v {
			if v2 == occupied {
				count++
			}
		}
	}
	return count
}

func copyBoard(curState [][]state) [][]state {
	out := make([][]state, len(curState))
	for k, v := range curState {
		out[k] = make([]state, len(v))
		copy(out[k], v)
	}
	return out
}

func equal(a, b [][]state) bool {
	for k, v := range a {
		for k2, v2 := range v {
			if v2 != b[k][k2] {
				return false
			}
		}
	}
	return true
}

func main() {
	process.ProcessData(&day11b{}, "data11.txt")
}
