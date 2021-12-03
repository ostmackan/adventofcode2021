package fileIo

import (
	"bufio"
	"os"
)

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
