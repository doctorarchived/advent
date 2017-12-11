package main

import (
	"io/ioutil"
	"fmt"
	"time"
)

func main()  {
	defer elapsed(time.Now())()
	PartTwo()
}
func PartOne()  {

	s, _ := ioutil.ReadFile("input.txt")
	sum := 0
	last := 0

	for i := 0; i < len(s) + 1; i++ {
		val := int(s[i % len(s)] - '0')
		if val == last {
			sum += val
		}

		last = val
	}

	fmt.Println(sum)
}
func PartTwo()  {
	s, _ := ioutil.ReadFile("input.txt")
	sum := 0

	for i := 0; i < len(s)/2; i++ {
		val1 := s[i]
		val2 := s[i + len(s)/2]
		if val1 == val2 {
			sum += 2* int(val1 - '0')
		}

	}

	fmt.Println(sum)
}

func elapsed(start time.Time) func() {
	return func() {
		fmt.Printf("Took %v\n",  time.Since(start))
	}
}

