package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"sort"
)

func main() {
	day4part2()

}

func day4part1()  {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	file, e := os.Open("input4.txt")
	fmt.Println(e)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		passphrase := scanner.Text()

		words := strings.Fields(passphrase)
		fmt.Println(words)
		sam := map[string]bool{}

		valid := true
		for _, w := range words {
			if sam[w] {
				valid = false
				break
			} else {
				sam[w] = true
			}
		}
		if valid {
			count++
		}
		fmt.Println(count)

	}
}

func day4part2() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	file, e := os.Open("input4.txt")
	fmt.Println(e)

	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		passphrase := scanner.Text()

		words := strings.Fields(passphrase)
		fmt.Println(words)
		sam := map[string]bool{}

		valid := true
		for _, w := range words {
			chars := strings.Split(w, "")
			sort.Strings(chars)
			sorted := strings.Join(chars, "")

			if sam[sorted] {
				valid = false
				break
			} else {
				sam[sorted] = true
			}
		}
		if valid {
			count++
		}
		fmt.Println(count)

	}

}
