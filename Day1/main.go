package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	body, err := os.ReadFile("E:\\sources\\advent-of-code-22-go\\Day1\\input.txt")
	if err != nil {
		fmt.Println("unable to read file: %v", err)
	}
	fmt.Println("Read sucessfully")
	lines := strings.Split(string(body), "\n")

	acctualSum, max := 0, 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if acctualSum > max {
				max = acctualSum
			}
			acctualSum = 0
		} else {
			acctualSum += tryParse(lines[i])
		}
	}
	fmt.Println("Max is: ", max)
}

func tryParse(value string) int {
	if val, err := strconv.ParseInt(value, 0, 0); err == nil {
		return int(val)
	} else {
		return 0
	}

}
