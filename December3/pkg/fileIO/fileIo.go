package fileIo

import (
	"bufio"
	"os"
	"strconv"
)

func GetIntegers(filename string) []int {

	inputFile, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	} else {
		scanner := bufio.NewScanner(inputFile)
		var values []int
		for scanner.Scan() {
			var val, err = strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err.Error())
			} else {
				values = append(values, val)
			}

		}
		return values
	}
}

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			var line = scanner.Text()
			lines = append(lines, line)
		}

		return lines
	}

}

func ConvertToIntArray(line string) []int {
	var values []int
	for i := range line {
		values = append(values, int(line[i]-'0'))
	}
	return values
}

func ReadAsIntegerArray(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		var values [][]int
		for scanner.Scan() {
			var line = scanner.Text()
			value := ConvertToIntArray(line)
			if err != nil {
				panic(err)
			} else {
				values = append(values, value)
			}
		}

		return values
	}
}
