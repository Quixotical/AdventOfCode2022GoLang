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
	rows := [][]string{}
	rows = append(rows, []string{"N", "H", "S", "J", "F", "W", "T", "D"})
	rows = append(rows, []string{"G", "B", "N", "T", "Q", "P", "R", "H"})
	rows = append(rows, []string{"V", "Q", "L"})
	rows = append(rows, []string{"Q", "R", "W", "S", "B", "N"})
	rows = append(rows, []string{"B", "M", "V", "T", "F", "D", "N"})
	rows = append(rows, []string{"R", "T", "H", "V", "B", "D", "M"})
	rows = append(rows, []string{"J", "Q", "B", "D"})
	rows = append(rows, []string{"Q", "H", "Z", "R", "V", "J", "N", "D"})
	rows = append(rows, []string{"S", "M", "H", "N", "B"})
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moveCommands := strings.Split(scanner.Text(), " ")
		numberToMove, err := strconv.Atoi(moveCommands[1])
		if err != nil {
			log.Fatal(err)
		}
		fromContainer, err := strconv.Atoi(moveCommands[3])
		if err != nil {
			log.Fatal(err)
		}
		toContainer, err := strconv.Atoi(moveCommands[5])
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range rows[fromContainer-1] {
			if k < numberToMove {
				rows[toContainer-1] = append([]string{v}, rows[toContainer-1]...)
			} else {
				rows[fromContainer-1] = rows[fromContainer-1][numberToMove:]
				break
			}
		}
	}
	for k, _ := range rows {
		fmt.Printf(rows[k][0])
	}
}
