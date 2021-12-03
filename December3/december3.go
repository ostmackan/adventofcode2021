package main

import (
	fileIo "aoc2021/december3/pkg/fileIo"
	"fmt"
	"strconv"
)

func main() {
	var (
		lines  [][]int
		result []int
		most   string = ""
		least  string = ""
	)
	lines = fileIo.ReadAsIntegerArray("input.txt")

	if len(lines) > 0 {
		for i := range lines[0] {
			var (
				ones   int = 0
				zeroes int = 0
			)
			for l := range lines {
				if lines[l][i] == 1 {
					ones++
				} else {
					zeroes++
				}
			}
			if ones > zeroes {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}

		}
		for i := range result {
			most += fmt.Sprint(result[i])
			if result[i] == 0 {
				least += fmt.Sprint(1)
			} else {
				least += fmt.Sprint(0)
			}

		}

		mostValue, err := strconv.ParseInt(most, 2, 16)
		if err != nil {
			panic(err)
		}
		leastValue, err := strconv.ParseInt(least, 2, 16)
		if err != nil {
			panic(err)
		}

		fmt.Printf("most: %s %d\n", most, mostValue)
		fmt.Printf("most: %s %d\n", least, leastValue)
		fmt.Println("answer: " + fmt.Sprint(mostValue*leastValue))
	}

}
