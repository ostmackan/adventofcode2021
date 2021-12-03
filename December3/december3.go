package main

import (
	fileIo "aoc2021/december3/pkg/fileIo"
	"fmt"
	"strconv"
)

func ConvertTo(line []int) string {
	var result string = ""
	for i := range line {
		result += fmt.Sprint(line[i])
	}
	return result
}

func MoreZeroesOrOnes(lines [][]int, i int) int {
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
	if ones >= zeroes {
		return 1
	} else {
		return 0
	}
}

func Keep(lines [][]int, i int, keep int) [][]int {
	var result [][]int

	for l := range lines {
		if lines[l][i] == keep {
			result = append(result, lines[l])
		}
	}

	return result
}

func main() {
	var (
		lines       [][]int
		mostLines   [][]int
		leastLines  [][]int
		result      []int
		most        string = ""
		least       string = ""
		mostResult  string
		leastResult string
	)
	lines = fileIo.ReadAsIntegerArray("input.txt")
	mostLines = fileIo.ReadAsIntegerArray("input.txt")
	leastLines = fileIo.ReadAsIntegerArray("input.txt")

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

			if len(mostLines) > 1 {
				mostLinesZeroOrOne := MoreZeroesOrOnes(mostLines, i)
				mostLines = Keep(mostLines, i, mostLinesZeroOrOne)
			}
			if len(leastLines) > 1 {
				leastLinesZeroOrOne := MoreZeroesOrOnes(leastLines, i)
				var leastResult = 0
				if leastLinesZeroOrOne == 0 {
					leastResult = 1
				}
				leastLines = Keep(leastLines, i, leastResult)
			}

			fmt.Printf("mostLines: %v\n", mostLines)
			fmt.Printf("leastLines: %v\n", leastLines)

		}
		for i := range result {
			most += fmt.Sprint(result[i])
			if result[i] == 0 {
				least += fmt.Sprint(1)
			} else {
				least += fmt.Sprint(0)
			}
		}

		if len(leastLines) == 1 {
			leastResult = ConvertTo(leastLines[0])
		} else {
			panic("leastLines more or less than 1")
		}

		if len(mostLines) == 1 {
			mostResult = ConvertTo(mostLines[0])
		} else {
			panic("leastLines more or less than 1")
		}

		mostValue, err := strconv.ParseInt(mostResult, 2, 16)
		if err != nil {
			panic(err)
		}
		leastValue, err := strconv.ParseInt(leastResult, 2, 16)
		if err != nil {
			panic(err)
		}

		fmt.Printf("most: %s %d\n", mostResult, mostValue)
		fmt.Printf("most: %s %d\n", leastResult, leastValue)
		fmt.Println("answer: " + fmt.Sprint(mostValue*leastValue))

	}

}
