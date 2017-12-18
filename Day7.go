package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"fmt"
	"strconv"
)

var r, _ = regexp.Compile("([a-z]+) \\(([0-9]+)\\)(?: -> ([a-z]+(?:, [a-z]+)+))?")

var child = make(map[string]string)

var parentWeights = make(map[string][]int)
var parents = make(map[string][]string)

var weights = make(map[string]int)

func main() {
	day7part1()
}

func day7part1() {

	file, _ := os.Open("input7.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		match := r.FindAllStringSubmatch(line, -1)

		disc := match[0][1]
		weight := match[0][2]
		onTop := match[0][3]

		discs := strings.Split(onTop, ", ")

		if _, ok := child[disc]; !ok {
			child[disc] = ""
		}

		for _, d := range discs {
			child[d] = disc
		}

		weights[disc], _ = strconv.Atoi(weight)
		parents[disc] = discs

	}

	var bottom string

	for k, v := range child {
		if v == "" {
			bottom = k
		}
	}
	recursionnnn(bottom)
	//fmt.Println()
	//fmt.Println(parents[bottom])
	//fmt.Println(parentWeights[bottom])

	fmt.Println(day7part2(bottom))

}

func day7part2(start string) int {

	cur := start
	diff := 0

	for cur != "" {
		next := ""

		posW, wizard := -1, -1
		posN, normie := -1, -1

		pw := parentWeights[cur]

		for i, w := range pw {
			//count[w] = append(count[w], parents[cur][i])
			if wizard == -1 && w != normie {
				wizard, posW = w, i
			} else if wizard == w {
				wizard, normie = normie, wizard
				posW, posN = posN, posW
			} else if normie == -1 {
				normie, posN = w, i
			}

		}

		if wizard == -1 {
			return diff + weights[cur]
		}

		if normie != -1 {
			next = parents[cur][posW]
			diff = normie - wizard
		}

		cur = next
	}

	return -1
}

func recursionnnn(disc string) int {

	parents := parents[disc]
	weight := weights[disc]

	var weights = make([]int, len(parents))

	for i, p := range parents {
		w := recursionnnn(p)
		weight += w
		weights[i] = w
	}

	parentWeights[disc] = weights

	return weight

}
