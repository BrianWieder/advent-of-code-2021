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

	scanner.Scan()
	l := scanner.Text()
	v1, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	l = scanner.Text()
	v2, err := strconv.Atoi(l)
	if err != nil {
		v2 = 0
		// log.Fatal(err)
	}
	scanner.Scan()
	l = scanner.Text()
	v3, err := strconv.Atoi(l)
	if err != nil {
		v3 = 0
		// log.Fatal(err)
	}

	res := 0
	for scanner.Scan() {
		l = scanner.Text()
		newValue, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		if v2+v3+newValue > v1+v2+v3 {
			res += 1
		}
		v1 = v2
		v2 = v3
		v3 = newValue
	}
	fmt.Println(res)
}
