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
	groupCounter := 0
	commonItems := map[rune]bool{}
	for scanner.Scan() {
		input := scanner.Text()
		log.Println(input)

		if groupCounter == 0 {
			for _, item := range input {
				commonItems[item] = true
			}
			groupCounter++
			continue
		}
		if groupCounter == 1 {
			newCommonItems := map[rune]bool{}
			for _, item := range input {
				if commonItems[item] {
					newCommonItems[item] = true
				}
			}
			commonItems = newCommonItems
			groupCounter++
			continue
		}

		for _, item := range input {
			if commonItems[item] {
				log.Println(string(item))
				sum += priorityValue(item)
				break
			}
		}
		commonItems = map[rune]bool{}
		groupCounter = 0
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
