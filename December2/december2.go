package main

import (
	fileIo "aoc2021/december2/pkg/fileIO"
	"fmt"
	"strconv"
	"strings"
)

type subCommand struct {
	direction string
	units     int
}

func GetCommands(lines []string) []subCommand {
	var subCommands []subCommand

	for i := range lines {
		s := strings.Split(lines[i], " ")
		var cmd subCommand
		if len(s) == 2 {
			cmd.direction = s[0]
			var units, err = strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			} else {
				cmd.units = units
				subCommands = append(subCommands, cmd)
			}
		}
	}

	return subCommands
}

func main() {
	var (
		lines    []string
		position int = 0
		depth    int = 0
	)
	lines = fileIo.ReadLines("input.txt")
	commands := GetCommands(lines)
	for i := range commands {
		command := commands[i]
		switch command.direction {
		case "forward":
			position += command.units
		case "up":
			depth -= command.units
		case "down":
			depth += command.units
		}
	}
	println("position: " + fmt.Sprint(position))
	println("depth: " + fmt.Sprint(depth))
	println("answer: " + fmt.Sprint(position*depth))
}
