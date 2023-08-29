package main

import (
	"fmt"
	"os"
	"strings"
)

type Round struct {
	played   HandShape
	response HandShape
}

func (round Round) getOutcome() int {
	if round.played == round.response {
		return 3
	} else if round.response == Scissors && round.played == Rock {
		return 0
	} else if round.response == Rock && round.played == Scissors {
		return 6
	} else if round.response > round.played {
		return 6
	} else {
		return 0
	}
}

func (round Round) getPoints() int {
	return round.getOutcome() + int(round.response)
}

func (round Round) getPlayedSign() HandShape {
	if round.response == Rock {
		return getLoose(round.played)
	} else if round.response == Paper { // draw
		return getDraw(round.played)
	} else { // win
		return getWin(round.played)
	}
}

func getLoose(played HandShape) HandShape {
	if played == Rock {
		return Scissors
	} else if played == Paper {
		return Rock
	} else {
		return Paper
	}
}

func getDraw(played HandShape) HandShape { return played }

func getWin(played HandShape) HandShape {
	if played == Rock {
		return Paper
	} else if played == Paper {
		return Scissors
	} else {
		return Rock
	}
}

func (round Round) getPointsFromRoundEnd() int {
	round.response = round.getPlayedSign()
	return round.getOutcome() + int(round.response)
}

type HandShape int

const (
	Rock     HandShape = 1
	Paper    HandShape = 2
	Scissors HandShape = 3
)

func main() {
	body, err := os.ReadFile("E:\\sources\\advent-of-code-22-go\\cmd\\Day2\\input.txt")
	if err != nil {
		fmt.Println("unable to read file: %v", err)
	}
	fmt.Println("Read sucessfully")
	round := parseMatch(string(body))
	total := 0
	for i := 0; i < len(round); i++ {
		total += round[i].getPointsFromRoundEnd()
	}
	fmt.Println(total)

}

func parseRound(val string) Round {
	parsed := strings.Split(val, " ")
	return Round{played: getShapeValue(parsed[0]), response: getShapeValue(parsed[1])}
}

func parseMatch(match string) []Round {
	roundValues := strings.Split(string(match), "\n")
	rounds := []Round{}
	for i := 0; i < len(roundValues); i++ {
		rounds = append(rounds, parseRound(roundValues[i]))
	}
	return rounds
}

func getShapeValue(shape string) HandShape {
	if shape == "A" || shape == "X" {
		return Rock
	} else if shape == "B" || shape == "Y" { //draw
		return Paper
	} else {
		return Scissors // win
	}
}
