package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commonItems := map[rune]bool{}
		input := scanner.Text()

		comptA := input[:(len(input)+1)/2]
		comptB := input[(len(input)+1)/2:]

		for _, item := range comptA {
			commonItems[item] = true
		}
		for _, item := range comptB {
			if commonItems[item] {
				sum += priorityValue(item)
				break
			}
		}
	}

	log.Println("answer:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func priorityValue(r rune) int {
	return strings.Index(chars, string(r)) + 1
}
