package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mostCalories := [3]int{0, 0, 0}
	totalElves := len(mostCalories)
	elfCalories := 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			mealCalories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalf("ERRER converting calories to text: %v", err)
			}
			elfCalories += mealCalories
		} else {
			// fmt.Println(elfCalories)
			if mostCalories[totalElves-1] > elfCalories {
				elfCalories = 0
				continue
			} else if mostCalories[totalElves-2] > elfCalories {
				mostCalories[totalElves-1] = elfCalories
			} else if mostCalories[totalElves-3] > elfCalories {
				mostCalories[totalElves-1] = mostCalories[totalElves-2]
				mostCalories[totalElves-2] = elfCalories
			} else {
				mostCalories[totalElves-1] = mostCalories[totalElves-2]
				mostCalories[totalElves-2] = mostCalories[totalElves-3]
				mostCalories[totalElves-3] = elfCalories
			}
			elfCalories = 0
		}
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	fmt.Println(mostCalories[0] + mostCalories[1] + mostCalories[2])
}
