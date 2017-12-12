package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"log"
	"strconv"
)

func main() {
	day5part1()
}

func day5part1() {

	file, e := os.Open("input6.txt")

	if e != nil {
		log.Println(e)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)

	blocks := make([]int, len(fields))

	for i := range fields {
		blocks[i], _ = strconv.Atoi(fields[i])
	}

	count := 0
	seen := make(map[string]int)

	seen[toString(blocks)] = count
	last := toString(blocks)
	for  {
		count++
		index, max := max(blocks)
		blocks[index] = 0
		for i := index + 1; i <= index+max; i++ {
			blocks[i%len(blocks)]++
		}

		last = toString(blocks)

		if _, ok := seen[last]; ok  {
			break
		} else {
			seen[last] = count
		}



	}

	fmt.Println(count) //part1
	fmt.Println(count - seen[last]) //part2

}

func max(guys []int) (int, int) {

	index, max := 0, 0
	for i, v := range guys {
		if v > max {
			index, max = i, v
		}
	}

	return index, max
}

func toString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", "", -1), "[]")

}
