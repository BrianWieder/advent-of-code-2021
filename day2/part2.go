package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var t string
	var distance int

	var aim = 0
	var h = 0
	var v = 0

	for n, err := fmt.Fscanf(file, "%s %d", &t, &distance); n != 0 && err == nil; n, err = fmt.Fscanf(file, "%s %d", &t, &distance) {
		switch t {
		case "forward":
			h += distance
			v += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}

	fmt.Println(h * v)

}
