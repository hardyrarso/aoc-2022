package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elfCalories := []int{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	total := 0
	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			elfCalories = append(elfCalories, total)
			total = 0
			continue
		}
		inputInt, _ := strconv.Atoi(input)
		total += inputInt
	}

	sort.Ints(elfCalories)
	log.Println("answer:", elfCalories[len(elfCalories)-3:])
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
