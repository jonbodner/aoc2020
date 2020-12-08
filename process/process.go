package process

import (
	"bufio"
	"bytes"
	"io/ioutil"
)

type Processor interface {
	Process(line string)
	Done()
}

func ProcessData(p Processor, filename string) {
	data, _ := ioutil.ReadFile(filename)
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		curRow := sc.Text()
		p.Process(curRow)
	}
	p.Done()
}

