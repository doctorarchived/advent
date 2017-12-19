package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
)

var r8, _ = regexp.Compile("([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) (<|>|<=|>=|==|!=) (-?[0-9]+)")

func main() {
	day8part1()

}


func day8part1() {

	file, _ := os.Open("input8.txt")

	scanner := bufio.NewScanner(file)
	mappy := make(map[string]int)
	allTimeMax := 0

	for scanner.Scan() {
		match := r8.FindAllStringSubmatch(scanner.Text(), 1)

		reg := match[0][1]
		op := match[0][2]
		val, _ := strconv.Atoi(match[0][3])

		regIf := match[0][4]
		opIf := match[0][5]
		valIf, _ := strconv.Atoi(match[0][6])


		if _, e := mappy[regIf]; !e {
			mappy[regIf] = 0
		}

		if _, e := mappy[reg]; !e {
			mappy[reg] = 0
		}

		if evalIf(mappy[regIf], opIf, valIf) {
			mappy[reg] = eval(mappy[reg], op, val)
			if mappy[reg] > allTimeMax {
				allTimeMax = mappy[reg]
			}
		}

	}

	max := 0

	for _,v := range mappy {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	fmt.Println(allTimeMax)

}

func eval(val1 int, op string, val2 int) int {

	if op == "inc" {
		return val1 + val2
	} else {
		return val1 - val2
	}
}

func evalIf(val1 int, op string, val2 int) bool {
	switch op {

	case "<":
		return val1 < val2
	case ">":
		return val1 > val2
	case "<=":
		return val1 <= val2
	case ">=":
		return val1 >= val2
	case "==":
		return val1 == val2
	case "!=":
		return val1 != val2
	default:
		return false

	}
}
