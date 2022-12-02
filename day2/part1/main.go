package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pointsMap := make(map[string]int)
	pointsMap["A X"] = 4
	pointsMap["A Y"] = 8
	pointsMap["A Z"] = 3
	pointsMap["B X"] = 1
	pointsMap["B Y"] = 5
	pointsMap["B Z"] = 9
	pointsMap["C X"] = 7
	pointsMap["C Y"] = 2
	pointsMap["C Z"] = 6
	comboMap := make(map[int]int)
	for scanner.Scan() {
		comboMap[pointsMap[scanner.Text()]] = comboMap[pointsMap[scanner.Text()]] + 1
	}
	points := 0
	for k := range comboMap {
		points += comboMap[k] * k
	}
	fmt.Printf("points: %v", points)

}
