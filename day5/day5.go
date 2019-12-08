package main

import (
	"adventofcode19/intcode"
	"adventofcode19/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please provide the path to your puzzle input: ")
	input, inputError := util.GetFileInput()
	if inputError != nil {
		fmt.Println(inputError.Error())
		return
	}

	var program []int
	sProgram := strings.Split(input, "\n")
	sProgram = strings.Split(sProgram[0], ",")
	for _, sInteger := range sProgram {
		if sInteger != "" {
			integer, convError := strconv.Atoi(sInteger)
			if convError != nil {
				fmt.Println("Failed to convert to integer:")
				fmt.Print(convError.Error())
			}
			program = append(program, integer)
		}
	}

	var intcodeInst intcode.Intcode
	programError := intcodeInst.RunProgram(&program)
	if programError != nil {
		fmt.Println(programError.Error())
		return
	}
}
