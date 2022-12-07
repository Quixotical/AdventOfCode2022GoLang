package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if str == v {
			return true
		}
	}
	return false
}

func main() {
	bufferSize := 14
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("ERROR reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	marker := []string{}
	markerRef := make(map[string]int)
	count := 0
	for scanner.Scan() {
		count++
		marker = append(marker, scanner.Text())
		markerRef[scanner.Text()] = markerRef[scanner.Text()] + 1
		if len(marker) >= bufferSize {
			if len(markerRef) == bufferSize {
				fmt.Println(count, markerRef)
				break
			}
			markerRef[marker[0]] = markerRef[marker[0]] - 1
			if markerRef[marker[0]] == 0 {
				delete(markerRef, marker[0])
			}
			marker = marker[1:]
		}
	}
}
