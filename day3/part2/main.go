package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	testCase := make(map[rune]int)
	count := 0
	total := 0
	for scanner.Scan() {
		for _, v := range scanner.Text() {
			if testCase[v] == count {
				testCase[v] = testCase[v] + 1
			}
		}
		if count == 2 {
			for k, v := range testCase {
				if v == 3 {
					amtToAdd := 0
					if k < 97 {
						amtToAdd = 26
					}
					total = total + (len(strings.Split(lowerCase, strings.ToLower(string(k)))[0]) + 1 + amtToAdd)
					break
				}
			}
			for k := range testCase {
				delete(testCase, k)
			}
			count = 0
		} else {
			count++
		}
	}
	fmt.Println(total)
}
