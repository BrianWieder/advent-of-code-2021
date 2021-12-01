package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	last := math.MaxUint32
	res := 0
	for scanner.Scan() {
		l := scanner.Text()
		newValue, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		if newValue > last {
			res += 1
		}
		last = newValue
	}
	fmt.Println(res)
}
