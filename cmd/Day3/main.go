package main

import (
	"fmt"
	"os"
	"strings"
)

type Pack struct {
	firstCompartment  string
	secondCompartment string
}

func (pack Pack) getItemsContinedInBothCompartment() []byte {
	values := make([]byte, 0)

	for i := 0; i < len(pack.firstCompartment); i++ {
		if strings.ContainsRune(pack.secondCompartment, rune(pack.firstCompartment[i])) &&
			!isAlreadyInArray(values, pack.firstCompartment[i]) {
			values = append(values, pack.firstCompartment[i])
		}
	}

	return values
}

func isAlreadyInArray(arr []byte, val byte) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

func parsePack(val string) Pack {
	len := len(val)
	return Pack{firstCompartment: val[:(len / 2)], secondCompartment: val[(len / 2):len]}
}

func parseInput(input string) []Pack {
	rows := strings.Split(string(input), "\n")
	packs := make([]Pack, 0)
	for i := 0; i < len(rows); i++ {
		packs = append(packs, parsePack(rows[i]))
	}
	return packs
}

func getValue(chars []byte) int {
	sum := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] >= 97 && chars[i] <= 122 {
			sum += int(chars[i]) - 96
		} else if chars[i] >= 65 && chars[i] <= 90 {
			sum += int(chars[i]) - 64 + 26
		}
	}
	return sum
}

func main() {
	body, err := os.ReadFile("E:\\sources\\advent-of-code-22-go\\cmd\\Day3\\input.txt")
	if err != nil {
		fmt.Println("unable to read file: %v", err)
		return
	}
	fmt.Println("Read sucessfully")
	packs := parseInput(string(body))
	val := 0
	for i := 0; i < len(packs); i++ {
		item := packs[i].getItemsContinedInBothCompartment()
		v := getValue(item)
		val += v
	}

	fmt.Println(val)
}
