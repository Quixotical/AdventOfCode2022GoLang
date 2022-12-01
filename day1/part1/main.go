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
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mostCalories := 0
	elfCalories := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			mealCalories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalf("ERRER converting calories to text: %v", err)
			}
			elfCalories += mealCalories
		} else {
			if elfCalories > mostCalories {
				mostCalories = elfCalories
			}
			elfCalories = 0
		}
	}
	fmt.Println(mostCalories)
}
