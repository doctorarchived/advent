package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"log"
	"strconv"
)

func main() {

	content, e := ioutil.ReadFile("input5.txt")
	if e != nil {
		log.Println(e)
	}

	stringOffsets := strings.Split(string(content), "\n")

	var intOffsets []int

	for _, s := range stringOffsets {
		i, e := strconv.Atoi(s)
		if e == nil {
			intOffsets = append(intOffsets, i)
		}
	}

	curIndex := 0
	steps := 0
	for curIndex >= 0 && curIndex < len(intOffsets){
		i := curIndex

		curIndex += intOffsets[curIndex]

		if intOffsets[i] >= 3 {
			intOffsets[i] -= 1
		} else {
			intOffsets[i] += 1
		}
		steps++

	}
	fmt.Println(steps)

}
