package main

import (
	"bufio"
	"fmt"
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

func slidingWindow(values []int, last int) int {
	if len(values) > last && last-2 >= 0 {
		return values[last-2] + values[last-1] + values[last]
	} else {
		panic("sW: " + fmt.Sprint(last) + " / " + fmt.Sprint(len(values)))
	}

}

func main() {
	integers := GetIntegers("input.txt")
	var increasedValues int = 0
	var lastValueSlidingWindow int = slidingWindow(integers, 2)
	for i := 2; i < len(integers); i++ {
		currentValue := slidingWindow(integers, i)
		if lastValueSlidingWindow < currentValue {
			increasedValues++
		}
		lastValueSlidingWindow = currentValue
	}

	fmt.Println(fmt.Sprint(increasedValues))
	fmt.Println(fmt.Sprint(lastValueSlidingWindow))
}
