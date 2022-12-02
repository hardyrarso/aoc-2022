package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Game struct {
	opponent, outcome string
}

var scoreTable = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var strategyTable = map[string]map[string]int{
	"X": {
		"A": 3,
		"B": 1,
		"C": 2,
	},
	"Y": {
		"A": 1,
		"B": 2,
		"C": 3,
	},
	"Z": {
		"A": 2,
		"B": 3,
		"C": 1,
	},
}

func (g *Game) Result() int {
	return scoreTable[g.outcome] + strategyTable[g.outcome][g.opponent]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		split := strings.Split(input, " ")

		game := &Game{split[0], split[1]}
		score += game.Result()
	}

	log.Println("answer:", score)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
