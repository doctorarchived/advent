package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	day9part1()
}



func day9part1() {



	file, _ := os.Open("input9.txt")
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanRunes)

	s.Split(bufio.ScanRunes)



	sum := 0
	depth := 0
	skip := false
	garbage := false

	cancelled := 0

	for s.Scan() {
		r := s.Text()

		if skip {
			skip = false
			continue
		}

		if !garbage {
			switch r {
			case "{":
				sum += depth + 1
				depth += 1
			case "}":
				depth += -1
			case "<":
				garbage = true
			}
		} else {
			switch r {
			case ">":
				garbage = false
			case "!":
				skip = true
			default:
				cancelled++
			}

		}
	}


	fmt.Println(sum)
	fmt.Println(cancelled)


}
