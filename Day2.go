package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main()  {


	part2()


}

func part1() {
	file, _ := os.Open("input2.txt")

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "\t")

		if len(split) > 0 {
			min,_ := strconv.Atoi(split[0])
			max,_ := strconv.Atoi(split[0])



			for _, v := range split {
				val,_ := strconv.Atoi(v)

				if val < min {
					min = val
				}

				if val > max {
					max = val
				}
			}

			diff := max - min

			sum += diff




		}



	}
	fmt.Println(sum)
}

func part2() {
	file, _ := os.Open("input2.txt")

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "\t")

		if len(split) > 0 {

			var history []int

			for _, v1 := range split {
				val,_ := strconv.Atoi(v1)

				for _, v2 := range history {
					if val % v2 == 0 {
						sum += val / v2
						fmt.Println(val, v2, val / v2)
					}

					if v2 % val == 0 {
						sum += v2 / val
						fmt.Println(v2, val, v2 / val)
					}
				}
				history = append(history, val)


			}

			fmt.Println(history)




		}



	}
	fmt.Println(sum)
}