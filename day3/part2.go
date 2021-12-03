package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
	}

	oxygen := ""
	oxygenLines := lines

	for i := 0; i < 12 && len(oxygen) == 0; i++ {
		zeroCount, oneCount := getCharacterCountInPos(oxygenLines, i)
		desiredBit := 0
		if zeroCount > oneCount {
			desiredBit = 0
		} else {
			desiredBit = 1
		}
		oxygenLines = filterLines(oxygenLines, i, desiredBit)
		if len(oxygenLines) == 1 {
			oxygen = oxygenLines[0]
		}
	}

	c02 := ""
	c02Lines := lines

	for i := 0; i < 12 && len(c02) == 0; i++ {
		zeroCount, oneCount := getCharacterCountInPos(c02Lines, i)
		desiredBit := 0
		if zeroCount <= oneCount {
			desiredBit = 0
		} else {
			desiredBit = 1
		}
		c02Lines = filterLines(c02Lines, i, desiredBit)
		if len(c02Lines) == 1 {
			c02 = c02Lines[0]
		}
	}

	o, err := strconv.ParseInt(oxygen, 2, 64)

	c, err := strconv.ParseInt(c02, 2, 64)

	fmt.Println(o * c)
}

func getCharacterCountInPos(lines []string, index int) (int, int) {
	zeroCount := 0
	oneCount := 0
	for _, s := range lines {
		if s[index] == '0' {
			zeroCount += 1
		} else {
			oneCount += 1
		}
	}
	return zeroCount, oneCount
}

func filterLines(lines []string, index int, desiredBit int) []string {
	newLines := make([]string, 0)

	for _, l := range lines {
		if l[index]-'0' == byte(desiredBit) {
			newLines = append(newLines, l)
		}
	}

	return newLines
}
