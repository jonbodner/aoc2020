package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	a := day4b{
		ps: []passport{{}},
	}
	processor(&a)
}

/*
byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
*/

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/

type passport map[string]string

var keys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

type day4a struct {
	ps []passport
}

/*
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm
*/
func (d *day4a) process(line string) {
	if len(line) == 0 {
		d.ps = append(d.ps, passport{})
		return
	}
	curP := d.ps[len(d.ps)-1]
	parts := strings.Split(line, " ")
	for _, v := range parts {
		kv := strings.Split(v, ":")
		curP[kv[0]] = kv[1]
	}
}

func (d *day4a) done() {
	count := 0
	for _, v := range d.ps {
		hasCID := false
		total := 0
		for _, k := range keys {
			if _, ok := v[k]; ok {
				total++
				if k == "cid" {
					hasCID = true
				}
			}
		}
		if (total == len(keys)-1 && !hasCID) || (total == len(keys)) {
			count++
		}
	}
	fmt.Println(count)
}

var keyRules = map[string]func(string) bool{
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	"byr": func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= 1920 && i <= 2002
	},
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	"iyr": func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= 2010 && i <= 2020
	},
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	"eyr": func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= 2020 && i <= 2030
	},
	// hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	"hgt": func(s string) bool {
		if strings.HasSuffix(s, "cm") {
			i, err := strconv.Atoi(s[:len(s)-2])
			if err != nil {
				return false
			}
			return i >= 150 && i <= 193
		}
		if strings.HasSuffix(s, "in") {
			i, err := strconv.Atoi(s[:len(s)-2])
			if err != nil {
				return false
			}
			return i >= 59 && i <= 76
		}
		return false
	},
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	"hcl": func(s string) bool {
		if s[0] != '#' {
			return false
		}
		if len(s) != 7 {
			return false
		}
		for _, v := range s[1:] {
			if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'f') {
				continue
			}
			return false
		}
		return true
	},
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	"ecl": func(s string) bool {
		valid := map[string]bool{
			"amb": true,
			"blu": true,
			"brn": true,
			"gry": true,
			"grn": true,
			"hzl": true,
			"oth": true,
		}
		return valid[s]
	},
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	"pid": func(s string) bool {
		if len(s) != 9 {
			return false
		}
		for _, v := range s {
			if v < '0' || v > '9' {
				return false
			}
		}
		return true
	},
	// cid (Country ID) - ignored, missing or not.
	"cid": func(s string) bool {
		return true
	},
}

type day4b struct {
	ps []passport
}

/*
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm
*/
func (d *day4b) process(line string) {
	if len(line) == 0 {
		d.ps = append(d.ps, passport{})
		return
	}
	curP := d.ps[len(d.ps)-1]
	parts := strings.Split(line, " ")
	for _, v := range parts {
		kv := strings.Split(v, ":")
		curP[kv[0]] = kv[1]
	}
}

func (d *day4b) done() {
	count := 0
	for _, v := range d.ps {
		hasCID := false
		total := 0
		for k, r := range keyRules {
			if val, ok := v[k]; ok && r(val) {
				total++
				if k == "cid" {
					hasCID = true
				}
			}
		}
		if (total == len(keys)-1 && !hasCID) || (total == len(keys)) {
			count++
		}
	}
	fmt.Println(count)
}

type Processor interface {
	process(line string)
	done()
}

func processor(p Processor) {
	data, _ := ioutil.ReadFile("data4.txt")
	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		curRow := sc.Text()
		p.process(curRow)
	}
	p.done()
}
