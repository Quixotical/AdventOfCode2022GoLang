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
	map1 := make(map[string]int)
	for scanner.Scan() {
		compartment1 := scanner.Text()[0 : len(scanner.Text())/2]
		compartment2 := scanner.Text()[len(scanner.Text())/2:]
		for _, c1v := range compartment1 {
			found := false
			for _, c2v := range compartment2 {
				if c1v == c2v {
					map1[string(c1v)] = map1[string(c1v)] + 1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	result := 0
	for _, v := range lowerCase {
		fmt.Println(v)
	}
	for _, v := range strings.ToUpper(lowerCase) {
		fmt.Println(v)
	}
	for k, v := range map1 {
		if strings.Contains(lowerCase, k) || strings.Contains(strings.ToUpper(lowerCase), k) {
			// fmt.Printf("key: %s, split list: %v\n", k, strings.Split(lowerCase, k)[0])
			amtToAdd := 0
			if []rune(k)[0] < 97 {
				amtToAdd = 26
			}
			// fmt.Printf("result is %v * letter(value): %s(%v) + %v\n", v, k, len(strings.SplitN(lowerCase, k, 2)[0])+1, amtToAdd)
			result = result + (v * (len(strings.SplitN(lowerCase, strings.ToLower(k), 2)[0]) + 1 + amtToAdd))
		}
	}
	fmt.Println(result)
}
