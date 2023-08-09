package main

import (
	"fmt"
	"os"
	"sort"
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

	acctualSum, maxes := 0, []int{0, 0, 0}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if isBiggerThanAny(maxes, acctualSum) {
				replaceLastAndSort(maxes, acctualSum)
			}
			acctualSum = 0
		} else {
			acctualSum += tryParse(lines[i])
		}
	}
	fmt.Println("Max is: ", sum(maxes))
}

func tryParse(value string) int {
	if val, err := strconv.ParseInt(value, 0, 0); err == nil {
		return int(val)
	} else {
		return 0
	}

}

func isBiggerThanAny(arr []int, val int) bool {
	for i := 0; i < len(arr); i++ {
		if val > arr[i] {
			return true
		}
	}
	return false
}

func replaceLastAndSort(arr []int, val int) {
	arr[0] = val
	sort.Ints(arr)
}

func sum(arr []int) int {
	sum := 0
	for _, valueInt := range arr {
		sum += valueInt
	}
	return sum
}
