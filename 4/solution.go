package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type cleanupSection struct {
	begin, end int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		split := strings.Split(input, ",")
		firstElf, secondElf := strings.Split(split[0], "-"), strings.Split(split[1], "-")

		firstElfBegin, _ := strconv.Atoi(firstElf[0])
		firstElfEnd, _ := strconv.Atoi(firstElf[1])
		firstElfSection := cleanupSection{
			begin: firstElfBegin,
			end:   firstElfEnd,
		}

		secondElfBegin, _ := strconv.Atoi(secondElf[0])
		secondElfEnd, _ := strconv.Atoi(secondElf[1])
		secondElfSection := cleanupSection{
			begin: secondElfBegin,
			end:   secondElfEnd,
		}

		log.Println("first:", firstElfSection)
		log.Println("second:", secondElfSection)
		if sectionOverlaps(firstElfSection, secondElfSection) || sectionOverlaps(secondElfSection, firstElfSection) {
			log.Println("includes")
			count++
		}
		//if count == 5 {
		//	break
		//}
	}

	log.Println("answer:", count)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sectionOverlaps(a, b cleanupSection) bool {
	return (a.end >= b.begin && a.end <= b.end) || (a.begin <= b.end && a.begin >= b.begin)
}
