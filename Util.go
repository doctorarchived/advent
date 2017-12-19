package main

import (
	"bufio"
	"os"
)

func scanner(filename string) *bufio.Scanner {

	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)

	return scanner
}