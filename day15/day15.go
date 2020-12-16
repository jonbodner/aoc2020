package main

import "fmt"

func main() {
	vals := []int{3, 1, 2}
	m := map[int][2]int{}
	for i, v := range vals {
		m[v] = [2]int{i, -1}
	}
	for i := len(vals) - 1; i < 30_000_000; i++ {
		pos, ok := m[vals[i]]
		if !ok {
			panic("should never happen!")
		}
		if pos[1] == -1 {
			vals = append(vals, 0)
			cur0, ok := m[0]
			if !ok {
				m[0] = [2]int{i + 1, -1}
			} else {
				m[0] = [2]int{i + 1, cur0[0]}
			}
			continue
		}
		newVal := i - pos[1]
		vals = append(vals, newVal)
		existing, ok := m[newVal]
		if !ok {
			m[newVal] = [2]int{i + 1, -1}
			continue
		}
		m[newVal] = [2]int{i + 1, existing[0]}
	}
	fmt.Println(vals[30_000_000-1])
}
