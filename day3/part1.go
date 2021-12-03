package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var counts [12][2]int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
		for i, c := range l {
			counts[i][c-'0'] += 1
		}
	}
	gamma := ""
	epsilon := ""
	for _, r := range counts {
		if r[0] > r[1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaRate, err := strconv.ParseInt(gamma, 2, 64)

	epsilonRate, err := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaRate * epsilonRate)
}
