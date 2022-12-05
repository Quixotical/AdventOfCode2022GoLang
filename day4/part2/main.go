package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	contains := 0
	for scanner.Scan() {
		lSide := strings.Split(strings.Split(scanner.Text(), ",")[0], "-")
		rSide := strings.Split(strings.Split(scanner.Text(), ",")[1], "-")
		r0, err := strconv.Atoi(rSide[0])
		if err != nil {
			log.Fatal(err)
		}
		r1, err := strconv.Atoi(rSide[1])
		if err != nil {
			log.Fatal(err)
		}
		l0, err := strconv.Atoi(lSide[0])
		if err != nil {
			log.Fatal(err)
		}
		l1, err := strconv.Atoi(lSide[1])
		if err != nil {
			log.Fatal(err)
		}
		if r0 >= l0 && r0 <= l1 {
			contains++
		} else if r1 <= l1 && r1 >= l0 {
			contains++
		} else if l0 >= r0 && l0 <= r1 {
			contains++
		} else if l1 <= r1 && l1 >= r0 {
			contains++
		}
	}
	fmt.Println(contains)
}
