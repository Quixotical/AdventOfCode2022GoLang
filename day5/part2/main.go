package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// [[J J F W T D] [N H S G B N T Q P R H] [V Q L] [Q R W S B N] [R B M V T F D N] [T H V B D M] [Q B D] [Q H Z R V J N D] [S M H N B]]
// [[J J F W T D] [N H S G B N T Q P R H] [V Q L] [Q R W S B N] [R B M V T S M H] [T H V B D M] [Q B D] [Q H Z R V J N D] [R B M V T S M H N B]]
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

		newFromSlice := rows[fromContainer-1][numberToMove:]
		newToSlice := rows[fromContainer-1][:numberToMove]
		rows[fromContainer-1] = nil
		for _, v := range newFromSlice {
			rows[fromContainer-1] = append(rows[fromContainer-1], v)
		}
		rows[toContainer-1] = append(newToSlice, rows[toContainer-1]...)
	}
	for k, _ := range rows {
		fmt.Printf("%s", rows[k][0])
	}
}
